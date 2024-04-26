package spider

import (
	"github.com/dcncy/gin-vue-admin/server/service"
)

type ApiGroup struct {
	SpiderTaskApi
}

var (
	spiderTaskService = service.ServiceGroupApp.DcncyServiceGroup.SpiderTaskService
)
