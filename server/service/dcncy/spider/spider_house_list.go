package spider

import (
	"fmt"
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/dcncy/gin-vue-admin/server/model/dcncy/house"
	"github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
	"github.com/dcncy/gin-vue-admin/server/utils/dcncy"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"time"
)

// 开始爬取
func (service *SpiderTaskService) StartSpiderTask(task *spider.SpiderTaskInfo) error {
	taskUrlPrefix := task.TaskUrlPrefix
	taskUrlSuffix := task.TaskUrlSuffix
	pageNum, _ := strconv.Atoi(task.PageNum)
	for i := 1; i <= pageNum; i++ {
		queryUrl := fmt.Sprintf("%s/pg%d%s", taskUrlPrefix, i, taskUrlSuffix)
		global.GVA_LOG.Info("=============>>>开始抓取网页：", zap.String("queryUrl", queryUrl))
		// 开始抓取
		err := spiderList(queryUrl)
		if err != nil {
			global.GVA_LOG.Error("抓取网页出现异常!", zap.String("queryUrl", queryUrl))
			return err
		}
	}
	return nil
}

// 创建爬虫对象
func spiderList(url string) error {
	// 创建 Collector 对象
	collector := CreateColly()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		// 添加请求头
		dcncy.CreateReqHeaderColly(request)
		global.GVA_LOG.Info("------请求调用之前:OnRequest------", zap.String("url", url))
	})
	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		global.GVA_LOG.Error("------请求错误:OnError------", zap.Error(err))
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		global.GVA_LOG.Info("------收到响应后调用:OnResponse------")
	})
	//OnResponse如果收到的内容是HTML ,则在之后调用
	mainElement := "div[class='dealListPage'] > div[class='content'] > div[class='leftContent'] > div[data-component='list'] > ul[class='listContent']"
	collector.OnHTML(mainElement, func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, el *colly.HTMLElement) {
			// 存储信息
			saveInfo(el)
		})
	})
	// 发起请求
	return collector.Visit(url)
}

// 保存爬取的内容
func saveInfo(el *colly.HTMLElement) {
	id := dcncy.GetSnowFlakeId()
	var aHouseOverview house.A_HouseOverview
	// 子爬虫使用该链接
	HouseDetailUrl := dcncy.TrimSpace(el.ChildAttr("a", "href"))
	// 贝壳主键ID
	IdBeike := dcncy.GetBeikeId(HouseDetailUrl)

	// 成交日期
	TradeDateStr := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='address'] > div[class='dealDate']"))
	// 校验是否已经爬取过
	existsKey := fmt.Sprintf("%s-%s", IdBeike, TradeDateStr)
	exists, err := checkExists(existsKey)
	if err != nil {
		global.GVA_LOG.Error("检查Redis是否缓存时出现异常!", zap.String("existsKey", existsKey), zap.Error(err))
	}
	// 如果没有抓取过，则才爬取
	if !exists {
		// 开启子协程
		go spiderHouseInfo(HouseDetailUrl, id)
		// 抓取列表数据
		err := parseHouseInfo(el, &aHouseOverview, id, IdBeike)
		if err != nil {
			global.GVA_LOG.Error("爬取列表数据出现异常!", zap.String("IdBeike", IdBeike), zap.Error(err))
		}
	} else {
		global.GVA_LOG.Info("该房源信息已经抓取过，本次不获取信息。", zap.String("IdBeike", IdBeike), zap.String("TradeDate", TradeDateStr))
	}
	// 每一房屋休眠三秒
	time.Sleep(global.SPIDER_CONFIG.Spider.SleepSecond * time.Second)
}

func parseHouseInfo(el *colly.HTMLElement, aHouseOverview *house.A_HouseOverview, id int, IdBeike string) error {
	// 户型图
	HouseLayoutSrc := dcncy.TrimSpace(el.ChildAttr("a > img[class='lj-lazy']", "data-original"))
	picPath, _ := dcncy.DownloadCoverPic(HouseLayoutSrc, id)
	// 房屋描述
	HouseTitle := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='title'] > a"))
	// 朝向
	Orientation := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='address'] > div[class='houseInfo']"))
	// 成交日期
	TradeDate := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='address'] > div[class='dealDate']"))
	// 成交总价
	TotalPrice := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='address'] > div[class='totalPrice'] > span[class='number']"))
	totalPrice, _ := strconv.ParseFloat(TotalPrice, 64)
	// 楼层信息
	FloorInfo := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='flood'] > div[class='positionInfo']"))
	// 每平单价
	UnitPrice := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='flood'] > div[class='unitPrice'] > span[class='number']"))
	unitPrice, _ := strconv.ParseFloat(UnitPrice, 64)
	// 房屋持有年限
	HoldYears := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='dealHouseInfo'] > span[class='dealHouseTxt']"))

	// 编译正则表达式，匹配数字部分
	re := regexp.MustCompile(`\d+`)

	// 挂牌价
	ListedPrice := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='dealCycleeInfo'] > span[class='dealCycleTxt'] > :nth-child(1)"))
	// 查找匹配的第一个数字
	listedPriceStr := re.FindString(ListedPrice)
	// 将匹配到的字符串转换为整数
	listedPrice, _ := strconv.ParseFloat(listedPriceStr, 64)
	// 成交周期
	TradeCycle := dcncy.TrimSpace(el.ChildText("div[class='info'] > div[class='dealCycleeInfo'] > span[class='dealCycleTxt'] > :nth-last-child(1)"))
	// 查找匹配的第一个数字
	tradeCycleStr := re.FindString(TradeCycle)
	tradeCycle, _ := strconv.Atoi(tradeCycleStr)
	// 组装入库
	aHouseOverview.Id = id
	aHouseOverview.IdBeike = IdBeike
	aHouseOverview.HouseTitle = HouseTitle
	aHouseOverview.HouseLayout = picPath
	aHouseOverview.Orientation = Orientation
	aHouseOverview.TradeDate = TradeDate
	aHouseOverview.TotalPrice = totalPrice
	aHouseOverview.FloorInfo = FloorInfo
	aHouseOverview.UnitPrice = unitPrice
	aHouseOverview.HoldYears = HoldYears
	aHouseOverview.ListedPrice = listedPrice
	aHouseOverview.TradeCycle = tradeCycle
	//存储数据
	err := global.GVA_DB.Create(&aHouseOverview).Error
	if err != nil {
		return err
	}
	// 存入缓存
	return dcncy.SPIDER_REDIS.Set(IdBeike, strconv.Itoa(id))
}

// 判断该房源是否已经抓取过
func checkExists(idBeike string) (bool, error) {
	var resultFlag = true
	// 1.判断redis缓存中是否存在
	_, flag := dcncy.SPIDER_REDIS.Get(idBeike, false)
	if !flag {
		// 2.如果缓存中没有，则查询数据库是否有
		var house house.A_HouseOverview
		result := global.GVA_DB.Where("id_beike = ?", idBeike).First(&house)
		if result.RowsAffected == 0 {
			// 3.数据库中也没有
			resultFlag = false
		} else {
			// 数据库中存在，存入换成
			err := dcncy.SPIDER_REDIS.Set(idBeike, strconv.Itoa(house.Id))
			if err != nil {
				return false, err
			}
		}
	}
	return resultFlag, nil
}

// 创建抓取工具
func CreateColly() colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("lf.ke.com", ".ke.com", ".ljcdn.com", "ke-image.ljcdn.com"), //白名单域名
		colly.UserAgent(global.SPIDER_CONFIG.Spider.SpiderHeader.Agent),
		//colly.AllowURLRevisit(), 						//允许对同一 URL 进行多次下载
		//colly.Async(true),                            //设置为异步请求
		//colly.MaxDepth(3),                            //爬取页面深度,最多为两层
		//colly.MaxBodySize(1024*1024*1024),            //响应正文最大字节数
		//colly.IgnoreRobotsTxt(), 						//忽略目标机器中的`robots.txt`声明
	)
	return *collector
}
