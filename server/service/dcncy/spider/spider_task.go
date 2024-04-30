package spider

import (
	"errors"
	"github.com/dcncy/gin-vue-admin/server/config"
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/dcncy/gin-vue-admin/server/model/common/request"
	"github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
	"github.com/dcncy/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type SpiderTaskService struct{}

//@author: [dcncy]
//@function: CreateTask
//@description: 新增爬虫任务
//@param: task spider.SpiderTaskInfo
//@return: spiderTaskInfo spider.SpiderTaskInfo, err error

func (service *SpiderTaskService) CreateTask(task spider.SpiderTaskInfo) (spiderTaskInfo spider.SpiderTaskInfo, err error) {
	var user spider.SpiderTaskInfo
	if !errors.Is(global.GVA_DB.Where("task_name = ?", task.TaskName).First(&user).Error, gorm.ErrRecordNotFound) {
		return spiderTaskInfo, errors.New("此爬虫任务名称已存在")
	}
	if !errors.Is(global.GVA_DB.Where("task_url_prefix = ? and task_url_suffix = ?", task.TaskUrlPrefix, task.TaskUrlSuffix).First(&user).Error, gorm.ErrRecordNotFound) {
		return spiderTaskInfo, errors.New("此爬虫任务链接已存在")
	}
	//存储数据
	err = global.GVA_DB.Create(&task).Error
	return task, err
}

//@author: [dcncy]
//@function: GetSpiderTaskInfoList
//@description: 分页获取数据
//@param: sysUserAuthorityID uint, info request.PageInfo
//@return: list interface{}, total int64, err error

func (service *SpiderTaskService) GetSpiderTaskInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spider.SpiderTaskInfo{})
	var CustomerList []spider.SpiderTaskInfo
	err = db.Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}

//@author: [dcncy]
//@function: DeleteSpiderTask
//@description: 删除爬虫任务
//@param: id float64
//@return: err error

func (service *SpiderTaskService) DeleteSpiderTask(id uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&spider.SpiderTaskInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [dcncy]
//@function: FindSpiderTaskById
//@description: 通过id获取爬虫任务信息
//@param: id int
//@return: err error, spiderTask *spider.SpiderTaskInfo

func (service *SpiderTaskService) FindSpiderTaskById(id uint) (spiderTask spider.SpiderTaskInfo, err error) {
	var task spider.SpiderTaskInfo
	err = global.GVA_DB.Where("id = ?", id).First(&task).Error
	return task, err
}

//@author: [dcncy]
//@function: UpdateSpiderTask
//@description: 更新爬虫任务
//@param: ID uint, status int8
//@return: err error

func (service *SpiderTaskService) UpdateSpiderTask(taskInfo spider.SpiderTaskInfo) (err error) {
	err = global.GVA_DB.Save(taskInfo).Error
	return err
}

//@author: [dcncy]
//@function: UpdateSpiderConfig
//@description: 更新爬虫相关配置

func (service *SpiderTaskService) UpdateSpiderConfig(conf config.SpiderConfig) (err error) {
	cs := utils.StructToMap(conf)
	for k, v := range cs {
		global.SPIDER_VP.Set(k, v)
	}
	err = global.SPIDER_VP.WriteConfig()
	return err
}

//@author: [dcncy]
//@function: UpdateSpiderConfig
//@description: 获取爬虫相关配置

func (service *SpiderTaskService) GetSpiderConfig() (conf config.Spider, err error) {
	return global.SPIDER_CONFIG.Spider, nil
}
