package spider

import (
	"github.com/dcncy/gin-vue-admin/server/service"
)

type ApiGroup struct {
	SpiderTaskApi
	SpiderDictionaryApi
	SpiderConfigApi
}

var (
	spiderTaskService       = service.ServiceGroupApp.DcncyServiceGroup.SpiderTaskService
	spiderDictionaryService = service.ServiceGroupApp.DcncyServiceGroup.SpiderDictionaryService
)
