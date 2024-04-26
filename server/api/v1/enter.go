package v1

import (
	"github.com/dcncy/gin-vue-admin/server/api/v1/dcncy/spider"
	"github.com/dcncy/gin-vue-admin/server/api/v1/example"
	"github.com/dcncy/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	SpiderApiGroup  spider.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
