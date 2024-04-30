package spider

import (
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/dcncy/gin-vue-admin/server/model/common/response"
	spiderModel "github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
	spiderReq "github.com/dcncy/gin-vue-admin/server/model/dcncy/spider/request"
	"github.com/dcncy/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SpiderDictionaryApi struct{}

//@author: [dcncy]
//@function: FindSpiderTaskById
//@description: 通过id获取爬虫任务信息

func (e *SpiderDictionaryApi) FindDictByType(c *gin.Context) {
	var dictionaryRequest spiderReq.SpiderDictionaryRequest
	err := c.ShouldBindJSON(&dictionaryRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(dictionaryRequest, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(dictionaryRequest, utils.SpiderDictionaryTypeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//
	taskInfoList, total, err := spiderDictionaryService.FindDictByType(dictionaryRequest)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     taskInfoList,
		Total:    total,
		Page:     dictionaryRequest.Page,
		PageSize: dictionaryRequest.PageSize,
	}, "获取成功", c)
}

//@author: [dcncy]
//@function: CreateSpiderDictionary
//@description: 创建爬虫字典项

func (e *SpiderDictionaryApi) CreateSpiderDictionary(c *gin.Context) {
	var spiderDictionary spiderModel.SpiderDictionary
	err := c.ShouldBindJSON(&spiderDictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(spiderDictionary, utils.SpiderDictionaryVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 存储数据
	_, err = spiderDictionaryService.CreateDictionary(spiderDictionary)
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

func (e *SpiderDictionaryApi) DeleteSpiderDictionary(c *gin.Context) {
	var dictionary spiderModel.SpiderDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(dictionary, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = spiderDictionaryService.DeleteSpiderDictionary(dictionary.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

//@author: [dcncy]
//@function: GetSpiderDictionary
//@description: 通过id获取爬虫字典

func (s *SpiderDictionaryApi) GetSpiderDictionary(c *gin.Context) {
	var dictionary spiderModel.SpiderDictionary
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(dictionary, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysDictionary, err := spiderDictionaryService.GetSpiderDictionary(dictionary.ID)
	if err != nil {
		global.GVA_LOG.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(sysDictionary, "查询成功", c)
}

//@author: [dcncy]
//@function: UpdateSpiderDictionary
//@description: 更新爬虫字典

func (e *SpiderDictionaryApi) UpdateSpiderDictionary(c *gin.Context) {
	var dictionary spiderModel.SpiderDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(dictionary, utils.SpiderDictionaryVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = spiderDictionaryService.UpdateSpiderDictionary(dictionary)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
