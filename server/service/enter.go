package service

import (
	"github.com/dcncy/gin-vue-admin/server/service/dcncy/spider"
	"github.com/dcncy/gin-vue-admin/server/service/example"
	"github.com/dcncy/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	DcncyServiceGroup   spider.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
