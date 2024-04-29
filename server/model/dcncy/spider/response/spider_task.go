package response

import (
	"github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
)

type SpiderTaskInfoResponse struct {
	SpiderTaskInfo spider.SpiderTaskInfo `json:"spiderTaskInfo"`
}
