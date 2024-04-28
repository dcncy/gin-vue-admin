package spider

import (
	"time"
)

const (
	Unprocessed int8 = iota
	Processing
	Completed
	Failed
	Drop
)

type SpiderTaskInfo struct {
	ID            uint       `json:"id" gorm:"primarykey"` // 主键ID
	TaskName      string     `json:"taskName" gorm:"comment:爬虫任务名称"`
	TaskUrlPrefix string     `json:"taskUrlPrefix" gorm:"comment:任务链接前缀"`
	TaskUrlSuffix string     `json:"taskUrlSuffix" gorm:"comment:任务链接后缀"`
	PageNum       string     `json:"pageNum" gorm:"default:1;comment:页数"`
	Status        string     `json:"status" gorm:"default:0;comment:处理状态: 0-未开始, 1-处理中, 2-处理完成, 3-处理失败, 4-废弃"`
	StartTime     *time.Time `json:"startTime" gorm:"default:null;comment:开始时间"`
	EndTime       *time.Time `json:"endTime" gorm:"default:null;comment:完成时间"`
}

func (SpiderTaskInfo) TableName() string {
	return "spider_task_info"
}
