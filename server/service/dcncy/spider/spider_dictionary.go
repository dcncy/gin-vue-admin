package spider

import (
	"errors"
	"github.com/dcncy/gin-vue-admin/server/global"
	"github.com/dcncy/gin-vue-admin/server/model/dcncy/spider"
	spiderReq "github.com/dcncy/gin-vue-admin/server/model/dcncy/spider/request"
	"gorm.io/gorm"
)

type SpiderDictionaryService struct{}

//@author: [dcncy]
//@function: FindDictByType
//@description: 分页获取数据

func (service *SpiderDictionaryService) FindDictByType(info spiderReq.SpiderDictionaryRequest) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spider.SpiderDictionary{}).Where("type = ?", info.Type).Order("id")
	var dataList []spider.SpiderDictionary
	err = db.Count(&total).Error
	if err != nil {
		return dataList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&dataList).Error
	}
	return dataList, total, err
}

//@author: [dcncy]
//@function: CreateTask
//@description: 创建爬虫字典项

func (service *SpiderDictionaryService) CreateDictionary(dictionary spider.SpiderDictionary) (spiderTaskInfo spider.SpiderDictionary, err error) {
	var dict spider.SpiderDictionary
	if !errors.Is(global.GVA_DB.Where("type = ? AND (name = ? OR code = ?)", dictionary.Type, dictionary.Name, dictionary.Code).First(&dict).Error, gorm.ErrRecordNotFound) {
		return spiderTaskInfo, errors.New("此类型下的字典已存在")
	}
	if !errors.Is(global.GVA_DB.Where("parent_code = ? AND (name = ? OR code = ?)", dictionary.ParentCode, dictionary.Name, dictionary.Code).First(&dict).Error, gorm.ErrRecordNotFound) {
		return spiderTaskInfo, errors.New("此父结点下的字典已存在")
	}
	//存储数据
	err = global.GVA_DB.Create(&dictionary).Error
	return dictionary, err
}

//@author: [dcncy]
//@function: DeleteSpiderTask
//@description: 删除爬虫任务
//@param: id float64
//@return: err error

func (service *SpiderDictionaryService) DeleteSpiderDictionary(id uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&spider.SpiderDictionary{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [dcncy]
//@function: GetSpiderDictionary
//@description: 通过id获取爬虫字典

func (service *SpiderDictionaryService) GetSpiderDictionary(id uint) (spiderTask spider.SpiderDictionary, err error) {
	var dictionary spider.SpiderDictionary
	err = global.GVA_DB.Where("id = ? and status = ?", id, spider.DICT_STATUS_ENABLED).First(&dictionary).Error
	return dictionary, err
}

//@author: [dcncy]
//@function: UpdateSpiderTask
//@description: 更新爬虫字典

func (service *SpiderDictionaryService) UpdateSpiderDictionary(dictionary spider.SpiderDictionary) (err error) {
	err = global.GVA_DB.Save(dictionary).Error
	return err
}
