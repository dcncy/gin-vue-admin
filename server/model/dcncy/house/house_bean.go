package house

// 房屋概览信息
type A_HouseOverview struct {
	Id          int     `json:"id" gorm:"index;comment:主键"`            // 主键
	IdBeike     string  `json:"id_beike" gorm:"index;comment:贝壳主键"`    // 贝壳主键
	HouseTitle  string  `json:"A_house_titile" gorm:"comment:房屋描述"`    // 房屋描述
	HouseLayout string  `json:"A_house_layout" gorm:"comment:户型图"`     // 户型图
	Orientation string  `json:"A_orientation" gorm:"comment:朝向"`       // 朝向
	TradeDate   string  `json:"A_trade_date" gorm:"comment:成交日期"`      // 成交日期
	TotalPrice  float64 `json:"A_total_price" gorm:"comment:成交总价(万元)"` // 成交总价(万元)
	FloorInfo   string  `json:"A_floor_info" gorm:"comment:楼层信息"`      // 楼层信息
	UnitPrice   float64 `json:"A_unit_price" gorm:"comment:每平单价(万/平)"` // 每平单价(万/平)
	HoldYears   string  `json:"A_hold_years" gorm:"comment:房屋持有年限"`    // 房屋持有年限
	ListedPrice float64 `json:"A_listed_price" gorm:"comment:挂牌价(万元)"` // 挂牌价(万元)
	TradeCycle  int     `json:"A_trade_cycle" gorm:"comment:成交周期(天)"`  // 成交周期(天)
}

func (A_HouseOverview) TableName() string {
	return "a_house_overview"
}

// 房屋交易信息
type B_HouseTradeInfo struct {
	HouseId        int     `json:"house_id" gorm:"index;comment:主键"`          // 主键
	CommunityId    int     `json:"B_community_id" gorm:"index;comment:小区主键"`  // 小区主键
	HouseLayout    string  `json:"B_house_layout" gorm:"comment:户型图详情"`       // 户型图详情
	TotalPrice     float64 `json:"B_total_price" gorm:"comment:成交总价(万元)"`     // 成交总价(万元)
	UnitPrice      float64 `json:"B_unit_price" gorm:"comment:每平单价(万/平)"`     // 每平单价(万/平)
	ListedPrice    float64 `json:"B_listed_price" gorm:"comment:挂牌价(万元)"`     // 挂牌价(万元)
	TradeCycle     int     `json:"B_trade_cycle" gorm:"comment:成交周期(天)"`      // 成交周期(天)
	AdjustPriceNum int     `json:"B_adjust_price_num" gorm:"comment:调价次数(次)"` // 调价次数(次)
	TakeViewTimes  int     `json:"B_take_view_times" gorm:"comment:带看次数(次)"`  // 带看次数(次)
	StarNum        int     `json:"B_star_num" gorm:"comment:关注人数(人)"`         // 关注人数(人)
	ViewNum        int     `json:"B_view_num" gorm:"comment:浏览次数(次)"`         // 浏览次数(次)
}

func (B_HouseTradeInfo) TableName() string {
	return "b_house_trade_info"
}

// 房屋基本信息
type C_HouseBaseInfo struct {
	HouseId           int     `json:"house_id" gorm:"index;comment:主键"`         // 主键
	HouseLayoutType   string  `json:"C_house_layout_type" gorm:"comment:房屋户型"`  // 房屋户型
	HouseFloor        string  `json:"C_house_floor" gorm:"comment:所在楼层"`        // 所在楼层
	TotalFloor        int     `json:"C_total_floor" gorm:"comment:总楼层"`         // 总楼层
	BuildingArea      float64 `json:"C_building_area" gorm:"comment:建筑面积(㎡)"`   // 建筑面积(㎡)
	IndoorArea        float64 `json:"C_indoor_area" gorm:"comment:套内面积(㎡)"`     // 套内面积(㎡)
	LayoutStructure   string  `json:"C_layout_structure" gorm:"comment:户型结构"`   // 户型结构
	BuildingType      string  `json:"C_building_type" gorm:"comment:建筑类型"`      // 建筑类型
	HouseOrientation  string  `json:"C_house_orientation" gorm:"comment:房屋朝向"`  // 房屋朝向
	ConstructionEra   int     `json:"C_construction_era" gorm:"comment:建成年代"`   // 建成年代
	RenovateInfo      string  `json:"C_renovate_info" gorm:"comment:装修情况"`      // 装修情况
	BuildingStructure string  `json:"C_building_structure" gorm:"comment:建筑结构"` // 建筑结构
	HeatingType       string  `json:"C_heating_type" gorm:"comment:供暖方式"`       // 供暖方式
	StaircaseRatio    string  `json:"C_staircase_ratio" gorm:"comment:梯户比例"`    // 梯户比例
	Elevator          string  `json:"C_elevator" gorm:"comment:配备电梯"`           // 配备电梯
	LianJiaSerialNo   string  `json:"C_lian_jia_serial_no" gorm:"comment:链家编号"` // 链家编号
	TradeOwnership    string  `json:"C_trade_ownership" gorm:"comment:交易权属"`    // 交易权属
	ListingDate       string  `json:"C_listing_date" gorm:"comment:挂牌时间"`       // 挂牌时间
	HousePurpose      string  `json:"C_house_purpose" gorm:"comment:房屋用途"`      // 房屋用途
	HoldYears         string  `json:"C_hold_years" gorm:"comment:房屋年限"`         // 房屋年限
	PropertyOwnership string  `json:"C_property_ownership" gorm:"comment:房权所属"` // 房权所属
}

func (C_HouseBaseInfo) TableName() string {
	return "c_house_base_info"
}

// 房屋历史成交记录
type D_HouseHistoryTrade struct {
	HouseId        int     `json:"house_id" gorm:"index;comment:主键"`              // 主键
	TotalPrice     float64 `json:"D_total_price" gorm:"comment:交易总价(万元)"`         // 交易总价(万元)
	TradeUnitPrice float64 `json:"d_trade_unit_price" gorm:"comment:历史交易单价(元/平)"` // 历史交易单价(元/平)
	TradeDate      string  `json:"d_trade_date" gorm:"comment:历史交易日期"`            // 历史交易日期
	TradeInfo      string  `json:"D_trade_info" gorm:"comment:交易详情"`              // 交易详情
}

func (D_HouseHistoryTrade) TableName() string {
	return "d_house_history_trade"
}

// 小区地址信息
type E_CommunityAddress struct {
	Id               int    `json:"id" gorm:"index;comment:主键"`            // 主键
	City             string `json:"E_city" gorm:"comment:城市"`              // 城市
	Region           string `json:"E_region" gorm:"comment:区域"`            // 区域
	BusinessDistrict string `json:"E_business_district" gorm:"comment:商圈"` // 商圈
	CommunityName    string `json:"E_community_name" gorm:"comment:小区"`    // 小区
}

func (E_CommunityAddress) TableName() string {
	return "e_community_address"
}
