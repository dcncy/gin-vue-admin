package dcncy

import (
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/gocolly/colly/v2"
)

func CreateReqHeaderColly(req *colly.Request) {
	req.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Headers.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Headers.Set("Connection", "keep-alive")
	req.Headers.Set("DNT", "1")
	req.Headers.Set("Host", "lf.ke.com")
	req.Headers.Set("Origin", "https://lf.ke.com/")
	req.Headers.Set("Referer", "https://lf.ke.com/")
	req.Headers.Set("Sec-Fetch-Dest", "document")
	req.Headers.Set("Sec-Fetch-Mode", "navigate")
	req.Headers.Set("Sec-Fetch-Site", "none")
	req.Headers.Set("Sec-Fetch-User", "?1")
	req.Headers.Set("Upgrade-Insecure-Requests", "1")
	req.Headers.Set("Sec-Ch-Ua", "\"Chromium\";v=\"124\", \"Google Chrome\";v=\"124\", \"Not-A.Brand\";v=\"99\"")
	req.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Headers.Set("Sec-Ch-Ua-Platform", "macOS")
	req.Headers.Set("Cookie", global.SPIDER_CONFIG.Spider.SpiderHeader.Cookie)
	req.Headers.Set("User-Agent", global.SPIDER_CONFIG.Spider.SpiderHeader.Agent)
}
