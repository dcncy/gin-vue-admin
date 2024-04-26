package request

import (
	"github.com/dcncy/gin-vue-admin/server/model/common/request"
	"github.com/dcncy/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
