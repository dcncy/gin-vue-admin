package spider

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	SpiderTaskApi
}

var (
	spiderTaskService = service.ServiceGroupApp.DcncyServiceGroup.SpiderTaskService
)
