package router

import (
	"github.com/dcncy/gin-vue-admin/server/router/dcncy"
	"github.com/dcncy/gin-vue-admin/server/router/example"
	"github.com/dcncy/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Dcncy   dcncy.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
