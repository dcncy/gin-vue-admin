package spider

const (
	DICT_STATUS_ENABLED int = iota
	DICT_STATUS_DISABLED
)

type SpiderDictionary struct {
	ID         uint   `json:"id" gorm:"primarykey"` // 主键ID
	Type       string `json:"type" gorm:"comment:字典类型"`
	Code       string `json:"code" gorm:"comment:字典值"`
	Name       string `json:"name" gorm:"comment:字典名称"`
	ParentCode string `json:"parentCode" gorm:"comment:父字典值"`
	Status     string `json:"status" gorm:"default:0;comment:状态: 0-启用, 1-禁用"`
}

func (SpiderDictionary) TableName() string {
	return "spider_dictionary"
}
