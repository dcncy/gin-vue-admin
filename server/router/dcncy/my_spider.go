package dcncy

import (
	v1 "github.com/dcncy/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type SpiderRouter struct{}

func (s *SpiderRouter) InitSpiderRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("spider")
	spiderTaskApi := v1.ApiGroupApp.SpiderApiGroup.SpiderTaskApi
	{
		customerRouter.POST("task", spiderTaskApi.CreateSpiderTask)          // 创建爬虫任务
		customerRouter.DELETE("task", spiderTaskApi.DeleteSpiderTask)        // 删除爬虫任务
		customerRouter.PUT("task", spiderTaskApi.UpdateSpiderTask)           // 更新爬虫任务
		customerRouter.POST("taskById", spiderTaskApi.FindSpiderTaskById)    // 根据Id获取爬虫任务
		customerRouter.POST("taskList", spiderTaskApi.GetSpiderTaskInfoList) // 获取爬虫任务列表

		customerRouter.POST("updateConfig", spiderTaskApi.UpdateSpiderConfig) // 更新爬虫相关配置
		customerRouter.POST("getConfig", spiderTaskApi.GetSpiderConfig)       // 获取爬虫相关配置
	}
}
