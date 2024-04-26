Drop table if exists a_house_overview;
CREATE TABLE IF NOT EXISTS a_house_overview
(
    id           Bigint PRIMARY KEY COMMENT '主键',
    id_beike     VARCHAR(100) COMMENT '贝壳主键',
    house_title  VARCHAR(100) COMMENT '房屋描述',
    house_layout VARCHAR(1000) COMMENT '户型图',
    orientation  VARCHAR(100) COMMENT '朝向',
    trade_date   VARCHAR(100) COMMENT '成交日期',
    total_price  DOUBLE COMMENT '成交总价(万元)',
    floor_info   VARCHAR(100) COMMENT '楼层信息',
    unit_price   DOUBLE COMMENT '每平单价(万/平)',
    hold_years   VARCHAR(100) COMMENT '房屋持有年限',
    listed_price DOUBLE COMMENT '挂牌价(万元)',
    trade_cycle  int COMMENT '成交周期(天)'
) COMMENT = '房屋概览信息表';

Drop table if exists b_house_trade_info;
CREATE TABLE IF NOT EXISTS b_house_trade_info
(
    house_id         Bigint PRIMARY KEY COMMENT '主键',
    community_id     Bigint COMMENT '小区主键',
    house_layout     VARCHAR(1000) COMMENT '户型图详情',
    total_price      DOUBLE COMMENT '成交总价',
    unit_price       DOUBLE COMMENT '每平单价',
    listed_price     DOUBLE COMMENT '挂牌价',
    trade_cycle      int COMMENT '成交周期',
    adjust_price_num int COMMENT '调价次数',
    take_view_times  int COMMENT '带看次数',
    star_num         int COMMENT '关注人数',
    view_num         int COMMENT '浏览次数'
) COMMENT = '房屋交易信息表';

Drop table if exists c_house_base_info;
CREATE TABLE IF NOT EXISTS c_house_base_info
(
    house_id           Bigint PRIMARY KEY COMMENT '主键',
    house_layout_type  VARCHAR(100) COMMENT '房屋户型',
    house_floor        VARCHAR(100) COMMENT '所在楼层',
    total_floor        int COMMENT '总楼层',
    building_area      DOUBLE COMMENT '建筑面积(㎡)',
    indoor_area        DOUBLE COMMENT '套内面积(㎡)',
    layout_structure   VARCHAR(100) COMMENT '户型结构',
    building_type      VARCHAR(100) COMMENT '建筑类型',
    house_orientation  VARCHAR(100) COMMENT '房屋朝向',
    construction_era   int COMMENT '建成年代',
    renovate_info      VARCHAR(100) COMMENT '装修情况',
    building_structure VARCHAR(100) COMMENT '建筑结构',
    heating_type       VARCHAR(100) COMMENT '供暖方式',
    staircase_ratio    VARCHAR(100) COMMENT '梯户比例',
    elevator           VARCHAR(100) COMMENT '配备电梯',
    lian_jia_serial_no VARCHAR(100) COMMENT '链家编号',
    trade_ownership    VARCHAR(100) COMMENT '交易权属',
    listing_date       VARCHAR(100) COMMENT '挂牌时间',
    house_purpose      VARCHAR(100) COMMENT '房屋用途',
    hold_years         VARCHAR(100) COMMENT '房屋年限',
    property_ownership VARCHAR(100) COMMENT '房权所属'
) COMMENT = '房屋基本信息表';

Drop table if exists d_house_history_trade;
CREATE TABLE IF NOT EXISTS d_house_history_trade
(
    house_id         Bigint COMMENT '主键',
    total_price      DOUBLE COMMENT '交易总价(万元)',
    trade_unit_price DOUBLE COMMENT '历史交易单价(元/平)',
    trade_date       VARCHAR(100) COMMENT '历史交易日期',
    trade_info       VARCHAR(1000) COMMENT '交易详情',
    PRIMARY KEY (house_id, trade_date)
) COMMENT = '房屋历史成交记录表';

Drop table if exists e_community_address;
CREATE TABLE IF NOT EXISTS e_community_address
(
    id                Bigint PRIMARY KEY COMMENT '主键',
    city              VARCHAR(100) COMMENT '城市',
    region            VARCHAR(1000) COMMENT '区域',
    business_district VARCHAR(1000) COMMENT '商圈',
    community_name    VARCHAR(1000) COMMENT '小区名称'
) COMMENT = '小区地址信息表';

CREATE TABLE IF NOT EXISTS spider_task_info
(
    id               INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    task_name        VARCHAR(255) UNIQUE COMMENT '爬虫任务名称',
    task_url_prefix VARCHAR(255) COMMENT '任务链接前缀',
    task_url_suffix VARCHAR(255) COMMENT '任务链接后缀',
    total_count      INT COMMENT '总房源量',
    page_num         INT COMMENT '页数',
    status           INT COMMENT '处理状态: 0-未开始, 1-处理中, 2-处理完成, 3-处理失败, 4-废弃',
    start_time       DATETIME COMMENT '开始时间',
    end_time         DATETIME COMMENT '完成时间'
) COMMENT = '爬虫任务表';
