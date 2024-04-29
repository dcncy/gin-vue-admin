package config

import "time"

type SpiderConfig struct {
	Spider Spider `mapstructure:"spider" json:"spider" yaml:"spider"`
}
type Spider struct {
	SleepSecond   time.Duration `mapstructure:"sleepSecond" json:"sleepSecond" yaml:"sleepSecond"` // 爬虫间隔休眠时间（秒）
	SpiderHeader  SpiderHeader  `mapstructure:"header" json:"header" yaml:"header"`                // 爬虫请求头
	SpiderPicPath SpiderPicPath `mapstructure:"picPath" json:"picPath" yaml:"picPath"`             // 爬虫图片存储地址
}

type SpiderHeader struct {
	Cookie string `mapstructure:"cookie" json:"cookie" yaml:"cookie"` // 请求头的cookie
	Agent  string `mapstructure:"agent" json:"agent" yaml:"agent"`    // 请求头的agent
}

type SpiderPicPath struct {
	Cover  string `mapstructure:"cover" json:"cover" yaml:"cover"`    // 封面图存放地址
	Layout string `mapstructure:"layout" json:"layout" yaml:"layout"` // 户型图存放地址
}
