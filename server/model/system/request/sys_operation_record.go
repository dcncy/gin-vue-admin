package request

import (
	"github.com/dcncy/gin-vue-admin/server/model/common/request"
	"github.com/dcncy/gin-vue-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
