package response11

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/dcncy/spider"
)

type SpiderTaskInfoResponse struct {
	SpiderTaskInfo spider.SpiderTaskInfo `json:"spiderTaskInfo"`
}
