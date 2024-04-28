package spider

import (
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/dcncy/gin-vue-admin/server/model/common/request"
	"github.com/dcncy/gin-vue-admin/server/model/common/response"
	spiderModel "github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
	spiderResp "github.com/dcncy/gin-vue-admin/server/model/dcncy/spider/response"
	"github.com/dcncy/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SpiderTaskApi struct{}

//@author: [dcncy]
//@function: GetSpiderTaskInfoList
//@description: 分页获取数据

func (e *SpiderTaskApi) GetSpiderTaskInfoList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	taskInfoList, total, err := spiderTaskService.GetSpiderTaskInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     taskInfoList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

//@author: [dcncy]
//@function: CreateSpiderTask
//@description: 新增爬虫任务

func (e *SpiderTaskApi) CreateSpiderTask(c *gin.Context) {
	var task spiderModel.SpiderTaskInfo
	err := c.ShouldBindJSON(&task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(task, utils.SpiderTaskVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 存储数据
	_, err = spiderTaskService.CreateTask(task)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

//@author: [dcncy]
//@function: DeleteSpiderTask
//@description: 删除爬虫任务

func (e *SpiderTaskApi) DeleteSpiderTask(c *gin.Context) {
	var taskInfo spiderModel.SpiderTaskInfo
	err := c.ShouldBindJSON(&taskInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(taskInfo.ID, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = spiderTaskService.DeleteSpiderTask(taskInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

//@author: [dcncy]
//@function: UpdateSpiderTask
//@description: 更新爬虫任务

func (e *SpiderTaskApi) UpdateSpiderTask(c *gin.Context) {
	var taskInfo spiderModel.SpiderTaskInfo
	err := c.ShouldBindJSON(&taskInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(taskInfo, utils.SpiderTaskVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = spiderTaskService.UpdateSpiderTask(taskInfo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

//@author: [dcncy]
//@function: FindSpiderTaskById
//@description: 通过id获取爬虫任务信息

func (e *SpiderTaskApi) FindSpiderTaskById(c *gin.Context) {
	var taskInfo spiderModel.SpiderTaskInfo
	err := c.ShouldBindQuery(&taskInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(taskInfo.ID, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := spiderTaskService.FindSpiderTaskById(taskInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(spiderResp.SpiderTaskInfoResponse{SpiderTaskInfo: data}, "获取成功", c)
}
