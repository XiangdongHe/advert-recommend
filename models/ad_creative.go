package models

import (
	"time"
)

// AdCreative 广告创意表
type AdCreative struct {
	CreativeID   int64     `gorm:"primaryKey;autoIncrement;column:creative_id" json:"creative_id"`
	PlanID       int64     `gorm:"not null;column:plan_id" json:"plan_id"`
	CreativeType int32     `gorm:"type:tinyint;not null;column:creative_type;comment:1=image,2=video,3=text" json:"creative_type"`
	MediaURL     string    `gorm:"type:varchar(512);not null;column:media_url" json:"media_url"`
	Title        string    `gorm:"type:varchar(256);column:title" json:"title"`
	Description  string    `gorm:"type:varchar(512);column:description" json:"description"`
	Status       int32     `gorm:"type:tinyint;not null;default:1;column:status" json:"status"`
	CreateTime   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_time" json:"update_time"`
}

// TableName 指定表名
func (AdCreative) TableName() string {
	return "t_ad_creative"
}
