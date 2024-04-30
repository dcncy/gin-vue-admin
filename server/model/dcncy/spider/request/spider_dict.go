package request

import (
	"github.com/dcncy/gin-vue-admin/server/model/common/request"
	"github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
)

type SpiderDictionaryRequest struct {
	request.PageInfo
	spider.SpiderDictionary
}
