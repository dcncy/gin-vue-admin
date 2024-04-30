package dcncy

import (
	v1 "github.com/dcncy/gin-vue-admin/server/api/v1"
	"github.com/dcncy/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SpiderRouter struct{}

func (s *SpiderRouter) InitSpiderRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("spider")
	customerRouterWithoutRecord := Router.Group("spider").Use(middleware.OperationRecord())
	spiderTaskApi := v1.ApiGroupApp.SpiderApiGroup.SpiderTaskApi
	spiderConfigApi := v1.ApiGroupApp.SpiderApiGroup.SpiderConfigApi
	spiderDictionaryApi := v1.ApiGroupApp.SpiderApiGroup.SpiderDictionaryApi
	{
		customerRouter.POST("task", spiderTaskApi.CreateSpiderTask)                       // 创建爬虫任务
		customerRouter.DELETE("task", spiderTaskApi.DeleteSpiderTask)                     // 删除爬虫任务
		customerRouter.PUT("task", spiderTaskApi.UpdateSpiderTask)                        // 更新爬虫任务
		customerRouterWithoutRecord.POST("taskById", spiderTaskApi.FindSpiderTaskById)    // 根据Id获取爬虫任务
		customerRouterWithoutRecord.POST("taskList", spiderTaskApi.GetSpiderTaskInfoList) // 获取爬虫任务列表
	}
	{
		customerRouter.POST("getConfig", spiderConfigApi.GetSpiderConfig)                    // 获取爬虫相关配置
		customerRouter.POST("start", spiderConfigApi.StartSpiderTask)                        // 启动爬虫任务
		customerRouterWithoutRecord.POST("updateConfig", spiderConfigApi.UpdateSpiderConfig) // 更新爬虫相关配置
	}
	{
		customerRouter.POST("dictionary", spiderDictionaryApi.CreateSpiderDictionary)          // 创建爬虫字典项
		customerRouter.DELETE("dictionary", spiderDictionaryApi.DeleteSpiderDictionary)        // 删除爬虫字典项
		customerRouter.PUT("dictionary", spiderDictionaryApi.UpdateSpiderDictionary)           // 更新爬虫字典项
		customerRouterWithoutRecord.POST("findDictByType", spiderDictionaryApi.FindDictByType) // 根据字典类型获取字典项
		customerRouterWithoutRecord.GET("dictionary", spiderDictionaryApi.GetSpiderDictionary) // 根据ID获取爬虫字典项
	}
}
