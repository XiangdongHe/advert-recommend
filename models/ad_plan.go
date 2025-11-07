package models

import (
	"time"
)

// AdPlan 广告计划表
type AdPlan struct {
	PlanID        int64     `gorm:"primaryKey;autoIncrement;column:plan_id" json:"plan_id"`
	Name          string    `gorm:"type:varchar(128);not null;column:name" json:"name"`
	Objective     string    `gorm:"type:varchar(64);not null;column:objective;comment:投放目标，click/download/conversion" json:"objective"`
	Budget        float64   `gorm:"type:decimal(16,2);not null;column:budget;comment:总预算" json:"budget"`
	BidPrice      string    `gorm:"type:varchar(64);not null;column:bid_price;comment:出价模式，CPC/CPM/CPA 0.5元/0.01元/5元" json:"bid_price"`
	TargetingRule string    `gorm:"type:json;not null;column:targeting_rule;comment:地域/年龄/性别/兴趣/设备等定向条件" json:"targeting_rule"`
	StartTime     time.Time `gorm:"type:datetime;not null;column:start_time" json:"start_time"`
	EndTime       time.Time `gorm:"type:datetime;not null;column:end_time" json:"end_time"`
	Status        int32     `gorm:"type:tinyint;not null;default:1;column:status;comment:1=active,0=paused,2=ended" json:"status"`
	CreateTime    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_time" json:"update_time"`

	// 不映射数据库，只用来承载查询结果
	Creatives []*AdCreative `gorm:"-" json:"creatives"`
}

// TableName 指定表名
func (AdPlan) TableName() string {
	return "t_ad_plan"
}
