package spider

import (
	"github.com/dcncy/gin-vue-admin/server/config"
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/dcncy/gin-vue-admin/server/model/common/response"
	spiderModel "github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
	spiderResp "github.com/dcncy/gin-vue-admin/server/model/dcncy/spider/response"
	"github.com/dcncy/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type SpiderConfigApi struct{}

//@author: [dcncy]
//@function: UpdateSpiderConfig
//@description: 更新爬虫相关配置

func (s *SpiderConfigApi) UpdateSpiderConfig(c *gin.Context) {
	var conf config.SpiderConfig
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = spiderTaskService.UpdateSpiderConfig(conf)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

//@author: [dcncy]
//@function: UpdateSpiderConfig
//@description: 获取爬虫相关配置

func (s *SpiderConfigApi) GetSpiderConfig(c *gin.Context) {
	conf, err := spiderTaskService.GetSpiderConfig()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(spiderResp.SpiderConfigResponse{Spider: conf}, "获取成功", c)
}

//@author: [dcncy]
//@function: StartSpiderTask
//@description: 执行爬虫任务

func (s *SpiderConfigApi) StartSpiderTask(c *gin.Context) {
	var taskInfo spiderModel.SpiderTaskInfo
	err := c.ShouldBindJSON(&taskInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(taskInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	taskInfo, err = spiderTaskService.FindSpiderTaskById(taskInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 启动爬虫任务
	go startSpider(c, taskInfo)
	response.OkWithMessage("启动成功", c)
}

// 开始抓取
func startSpider(c *gin.Context, task spiderModel.SpiderTaskInfo) {
	// 开始爬取，更新状态
	now := time.Now()
	task.StartTime = &now
	task.EndTime = task.StartTime
	task.Status = strconv.Itoa(spiderModel.Processing)

	err := spiderTaskService.UpdateSpiderTask(task)
	if err != nil {
		global.GVA_LOG.Error("更新爬虫任务状态失败!", zap.Error(err))
		response.FailWithMessage("更新爬虫任务状态失败", c)
		return
	}
	// 启动爬虫任务
	err = spiderTaskService.StartSpiderTask(&task)
	if err != nil {
		global.GVA_LOG.Error("任务执行失败!", zap.Error(err))
		response.FailWithMessage("任务执行失败", c)
		// 更新爬虫任务为失败
		task.Status = strconv.Itoa(spiderModel.Failed)
	} else {
		task.Status = strconv.Itoa(spiderModel.Completed)
	}
	// 完成爬取，更新状态
	now = time.Now()
	task.EndTime = &now
	err = spiderTaskService.UpdateSpiderTask(task)
	if err != nil {
		global.GVA_LOG.Error("更新爬虫任务状态失败!", zap.Error(err))
		response.FailWithMessage("更新爬虫任务状态失败", c)
		return
	}
}
