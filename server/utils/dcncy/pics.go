package dcncy

import (
	"fmt"
	"github.com/dcncy/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func DownloadPic(imageURL string, id int) (string, error) {
	PicPath := global.SPIDER_CONFIG.Spider.SpiderPicPath.Cover
	// 使用filepath.Ext()函数获取文件后缀名
	fileExtTemp := strings.Split(imageURL, "?")[0]
	fileExt := filepath.Ext(fileExtTemp)

	// 要保存的文件路径，相对或绝对路径均可
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	// 生成1到100之间的随机数
	randomNumber := rand.Intn(100) + 1
	// 格式化输出，确保输出为三位数（不足三位则补零）
	randomNumberString := fmt.Sprintf("%03d", randomNumber)
	downloadPath := PicPath + "/" + strconv.Itoa(id) + "_" + randomNumberString + fileExt
	// 发起HTTP GET请求获取图片数据
	response, err := http.Get(imageURL)
	if err != nil {
		global.GVA_LOG.Error("无法下载图片!", zap.String("imageURL", imageURL), zap.Error(err))
		return "", err
	}
	defer response.Body.Close()
	// 创建要保存的目录
	err = os.MkdirAll(filepath.Dir(PicPath), os.ModePerm)
	if err != nil {
		global.GVA_LOG.Error("无法创建目录!", zap.String("PicPath", PicPath), zap.Error(err))
		return "", err
	}

	// 创建一个新文件用于保存图片
	file, err := os.Create(downloadPath)
	if err != nil {
		global.GVA_LOG.Error("无法创建文件!", zap.String("downloadPath", downloadPath), zap.Error(err))
		return "", err
	}
	defer file.Close()

	// 将图片数据写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		global.GVA_LOG.Error("无法保存图片!", zap.String("pic", downloadPath), zap.Error(err))
		return "", err
	}
	global.GVA_LOG.Info("图片下载并保存成功")
	return downloadPath, nil
}
