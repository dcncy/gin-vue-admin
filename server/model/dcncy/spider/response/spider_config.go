package response

import "github.com/dcncy/gin-vue-admin/server/config"

type SpiderConfigResponse struct {
	Spider config.Spider `json:"spider"`
}
