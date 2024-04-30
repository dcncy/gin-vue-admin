package spider

import (
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/dcncy/gin-vue-admin/server/model/dcncy/house"
	"github.com/dcncy/gin-vue-admin/server/utils/dcncy"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
)

func spiderHouseInfo(url string, id int) {
	global.GVA_LOG.Info("=============>>>开始子爬虫抓取网页：", zap.String("url", url))
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

	communityId := dcncy.GetSnowFlakeId()
	// E-房屋小区信息
	collector.OnHTML("div[data-component='detailHeader'] > div[mod-id='lj-common-bread'] > div[class='container']", func(element *colly.HTMLElement) {
		element.ForEach("div[class='fl l-txt']", func(i int, el *colly.HTMLElement) {
			// 存储信息
			communityId = save_E_CommunityAddress(el, communityId)
		})
	})
	// B-房屋交易信息
	collector.OnHTML("div[data-component='overviewIntro']", func(element *colly.HTMLElement) {
		element.ForEach("div[class='overview']", func(i int, el *colly.HTMLElement) {
			// 存储信息
			save_B_HouseTradeInfo(el, id, communityId)
		})
	})
	// C-房屋基本信息
	collector.OnHTML("div[data-component='baseinfo']", func(element *colly.HTMLElement) {
		element.ForEach("div[class='newwrap baseinform'] > div[class='introContent']", func(i int, el *colly.HTMLElement) {
			// 存储信息
			save_C_HouseBaseInfo(el, id)
		})
	})
	// D-房屋历史成交记录
	collector.OnHTML("div[data-component='dealRecord']", func(element *colly.HTMLElement) {
		element.ForEach("div[class='chengjiao_record'] > ul[class='record_list'] > li", func(i int, el *colly.HTMLElement) {
			// 存储信息
			save_D_HouseHistoryTrade(el, id)
		})
	})
	// 发起请求
	err := collector.Visit(url)
	if err != nil {
		return
	}
}

// 存储房屋交易信息
func save_B_HouseTradeInfo(el *colly.HTMLElement, id int, communityId int) {
	// 小户型图
	houseLayoutSrc := dcncy.TrimSpace(el.ChildAttr("div[class='img'] > div[class='thumbnail'] > ul > li", "data-src"))
	// 大户型图
	bigLayoutPic, exists := el.DOM.Parent().Find("div[class='bigImg'] > div[class='slide'] > ul > li").Attr("data-src")
	if exists {
		houseLayoutSrc = dcncy.TrimSpace(bigLayoutPic)
	}
	// 下载户型图
	picPath, _ := dcncy.DownloadLayoutPic(houseLayoutSrc, id)
	// 成交总价
	TotalPrice := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='price'] > span[class='dealTotalPrice'] > i"))
	totalPrice, _ := strconv.ParseFloat(TotalPrice, 64)
	// 每平单价
	UnitPrice := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='price'] > b"))
	unitPrice, _ := strconv.ParseFloat(UnitPrice, 64)
	// 挂牌价
	ListedPrice := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='msg'] > :nth-child(1) > label"))
	listedPrice, _ := strconv.ParseFloat(ListedPrice, 64)
	// 成交周期
	TradeCycle := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='msg'] > :nth-child(2) > label"))
	tradeCycle, _ := strconv.Atoi(TradeCycle)
	// 调价次数
	AdjustPriceNum := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='msg'] > :nth-child(3) > label"))
	adjustPriceNum, _ := strconv.Atoi(AdjustPriceNum)
	// 带看次数
	TakeViewTimes := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='msg'] > :nth-child(4) > label"))
	takeViewTimes, _ := strconv.Atoi(TakeViewTimes)
	// 关注人数
	StarNum := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='msg'] > :nth-child(5) > label"))
	starNum, _ := strconv.Atoi(StarNum)
	// 浏览次数
	ViewNum := dcncy.TrimSpace(el.ChildText("div[class='info fr'] > div[class='msg'] > :nth-child(6) > label"))
	viewNum, _ := strconv.Atoi(ViewNum)
	//组装数据
	var houseTradeInfo house.B_HouseTradeInfo
	// 主键
	houseTradeInfo.HouseId = id
	// 小区主键
	houseTradeInfo.CommunityId = communityId
	// 户型图
	houseTradeInfo.HouseLayout = picPath
	// 成交总价
	houseTradeInfo.TotalPrice = totalPrice
	// 每平单价
	houseTradeInfo.UnitPrice = unitPrice
	// 挂牌价
	houseTradeInfo.ListedPrice = listedPrice
	// 成交周期
	houseTradeInfo.TradeCycle = tradeCycle
	// 调价次数
	houseTradeInfo.AdjustPriceNum = adjustPriceNum
	// 带看次数
	houseTradeInfo.TakeViewTimes = takeViewTimes
	// 关注人数
	houseTradeInfo.StarNum = starNum
	// 浏览次数
	houseTradeInfo.ViewNum = viewNum
	// 入库
	global.GVA_DB.Create(&houseTradeInfo)
}

func save_C_HouseBaseInfo(el *colly.HTMLElement, id int) {
	// 1.房屋户型
	houseLayoutType := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(1)"))
	houseLayoutType = strings.TrimPrefix(houseLayoutType, "房屋户型")
	// 2.所在楼层
	houseFloor := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(2)"))
	houseFloor = strings.TrimPrefix(houseFloor, "所在楼层")
	// 总楼层
	// 定义正则表达式，匹配括号中的数字
	re := regexp.MustCompile(`\d+`)
	// 查找匹配的子串
	numberStr := re.FindString(houseFloor)
	totalFloor, _ := strconv.Atoi(numberStr)
	// 3.建筑面积(㎡)
	BuildingArea := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(3)"))
	BuildingArea = strings.TrimPrefix(BuildingArea, "建筑面积")
	BuildingArea = strings.TrimSuffix(BuildingArea, "㎡")
	buildingArea, _ := strconv.ParseFloat(BuildingArea, 64)
	// 套内面积(㎡)
	IndoorArea := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(5)"))
	IndoorArea = strings.TrimPrefix(IndoorArea, "套内面积")
	IndoorArea = strings.TrimSuffix(IndoorArea, "㎡")
	indoorArea, _ := strconv.ParseFloat(IndoorArea, 64)
	// 户型结构
	layoutStructure := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(4)"))
	layoutStructure = strings.TrimPrefix(layoutStructure, "户型结构")
	// 建筑类型
	buildingType := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(6)"))
	buildingType = strings.TrimPrefix(buildingType, "建筑类型")
	// 房屋朝向
	houseOrientation := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(7)"))
	houseOrientation = strings.TrimPrefix(houseOrientation, "房屋朝向")
	// 建成年代
	ConstructionEra := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(8)"))
	ConstructionEra = strings.TrimPrefix(ConstructionEra, "建成年代")
	constructionEra, _ := strconv.Atoi(ConstructionEra)
	// 装修情况
	renovateInfo := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(9)"))
	renovateInfo = strings.TrimPrefix(renovateInfo, "装修情况")
	// 建筑结构
	buildingStructure := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(10)"))
	buildingStructure = strings.TrimPrefix(buildingStructure, "建筑结构")
	// 供暖方式
	heatingType := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(11)"))
	heatingType = strings.TrimPrefix(heatingType, "供暖方式")
	// 梯户比例
	staircaseRatio := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(12)"))
	staircaseRatio = strings.TrimPrefix(staircaseRatio, "梯户比例")
	// 配备电梯
	elevator := dcncy.TrimSpace(el.ChildText("div[class='base'] > div[class='content'] > ul > :nth-child(13)"))
	elevator = strings.TrimPrefix(elevator, "配备电梯")
	// 链家编号
	lianJiaSerialNo := dcncy.TrimSpace(el.ChildText("div[class='transaction'] > div[class='content'] > ul > :nth-child(1)"))
	lianJiaSerialNo = strings.TrimPrefix(lianJiaSerialNo, "链家编号")
	// 交易权属
	tradeOwnership := dcncy.TrimSpace(el.ChildText("div[class='transaction'] > div[class='content'] > ul > :nth-child(2)"))
	tradeOwnership = strings.TrimPrefix(tradeOwnership, "交易权属")
	// 挂牌时间
	listingDate := dcncy.TrimSpace(el.ChildText("div[class='transaction'] > div[class='content'] > ul > :nth-child(3)"))
	listingDate = strings.TrimPrefix(listingDate, "挂牌时间")
	// 房屋用途
	housePurpose := dcncy.TrimSpace(el.ChildText("div[class='transaction'] > div[class='content'] > ul > :nth-child(4)"))
	housePurpose = strings.TrimPrefix(housePurpose, "房屋用途")
	// 房屋年限
	holdYears := dcncy.TrimSpace(el.ChildText("div[class='transaction'] > div[class='content'] > ul > :nth-child(5)"))
	holdYears = strings.TrimPrefix(holdYears, "房屋年限")
	// 房权所属
	propertyOwnership := dcncy.TrimSpace(el.ChildText("div[class='transaction'] > div[class='content'] > ul > :nth-child(6)"))
	propertyOwnership = strings.TrimPrefix(propertyOwnership, "房权所属")
	//组装数据
	var houseBaseInfo house.C_HouseBaseInfo
	// 主键
	houseBaseInfo.HouseId = id
	houseBaseInfo.HouseLayoutType = houseLayoutType
	houseBaseInfo.HouseFloor = houseFloor
	houseBaseInfo.TotalFloor = totalFloor
	houseBaseInfo.BuildingArea = buildingArea
	houseBaseInfo.IndoorArea = indoorArea
	houseBaseInfo.LayoutStructure = layoutStructure
	houseBaseInfo.BuildingType = buildingType
	houseBaseInfo.HouseOrientation = houseOrientation
	houseBaseInfo.ConstructionEra = constructionEra
	houseBaseInfo.RenovateInfo = renovateInfo
	houseBaseInfo.BuildingStructure = buildingStructure
	houseBaseInfo.HeatingType = heatingType
	houseBaseInfo.StaircaseRatio = staircaseRatio
	houseBaseInfo.Elevator = elevator
	houseBaseInfo.LianJiaSerialNo = lianJiaSerialNo
	houseBaseInfo.TradeOwnership = tradeOwnership
	houseBaseInfo.ListingDate = listingDate
	houseBaseInfo.HousePurpose = housePurpose
	houseBaseInfo.HoldYears = holdYears
	houseBaseInfo.PropertyOwnership = propertyOwnership
	// 入库
	global.GVA_DB.Create(&houseBaseInfo)
}

func save_D_HouseHistoryTrade(el *colly.HTMLElement, id int) {
	// 交易总价
	priceTemp := dcncy.TrimSpace(el.ChildText("span[class='record_price']"))
	priceTemp = strings.ReplaceAll(priceTemp, "万", "")
	totalPrice, _ := strconv.ParseFloat(priceTemp, 64)
	// 交易详情
	tradeInfo := dcncy.TrimSpace(el.ChildText("p[class='record_detail']"))
	parts := strings.Split(tradeInfo, ",")
	// 提取单价部分
	pricePart := parts[0]
	tradeUnitPriceTemp := strings.TrimPrefix(pricePart, "单价")
	tradeUnitPriceTemp = strings.TrimSuffix(tradeUnitPriceTemp, "元/平")
	tradeUnitPrice, _ := strconv.ParseFloat(tradeUnitPriceTemp, 64)
	// 提取日期部分
	TradeDate := parts[1]
	// 组装数据
	var houseHistoryTrade house.D_HouseHistoryTrade
	// 房屋主键
	houseHistoryTrade.HouseId = id
	// 交易总价
	houseHistoryTrade.TotalPrice = totalPrice
	// 交易详情
	houseHistoryTrade.TradeInfo = tradeInfo
	// 历史交易单价(元/平)
	houseHistoryTrade.TradeUnitPrice = tradeUnitPrice
	// // 历史交易日期
	houseHistoryTrade.TradeDate = TradeDate
	// 入库
	global.GVA_DB.Create(&houseHistoryTrade)
}

// 小区地址信息
func save_E_CommunityAddress(el *colly.HTMLElement, id int) int {
	// 城市
	city := dcncy.SimpleStrErShouFang(el.ChildText("a:nth-child(3)"))
	// 区域
	region := dcncy.SimpleStrErShouFang(el.ChildText("a:nth-child(5)"))
	// 商圈
	businessDistrict := dcncy.SimpleStrErShouFang(el.ChildText("a:nth-child(7)"))
	// 小区
	communityName := dcncy.SimpleStrErShouFang(el.ChildText("a:nth-child(9)"))
	// 1.判断缓存中是否有这个小区信息
	strSlice := []string{city, region, businessDistrict, communityName}
	cacheKey := strings.Join(strSlice, "-")
	cacheId, flag := dcncy.SPIDER_REDIS.Get(cacheKey, false)
	// 如果缓存中不存在
	if !flag {
		// 2.如果缓存中没有，则查询数据库是否有
		var community house.E_CommunityAddress
		result := global.GVA_DB.Where("city = ? AND region = ? AND business_district = ? AND community_name = ?", city, region, businessDistrict, communityName).First(&community)
		if result.RowsAffected == 0 {
			// 3.数据库中也没有的话，根据传进来的id，创建小区信息，并入库并入缓存
			community.Id = id
			community.City = city
			community.Region = region
			community.BusinessDistrict = businessDistrict
			community.CommunityName = communityName
			// 入库
			global.GVA_DB.Create(&community)
		} else {
			// 数据库中存在
			id = community.Id
		}
		// 存储到缓存
		dcncy.SPIDER_REDIS.Set(cacheKey, strconv.Itoa(id))
	} else {
		// 缓存中存在
		id, _ = strconv.Atoi(cacheId)
	}
	return id
}
