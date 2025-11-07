package models

import (
	"time"
)

// UserProfileInterest 用户兴趣画像（标签 + 权重）
type UserProfileInterest struct {
	ID         int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID     int64     `gorm:"not null;column:user_id;index:idx_user_tag,priority:1" json:"user_id"`
	Tag        string    `gorm:"type:varchar(128);not null;column:tag;index:idx_user_tag,priority:2" json:"tag"`
	Weight     float64   `gorm:"type:decimal(5,4);not null;column:weight;comment:0~1" json:"weight"`
	UpdateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_time" json:"update_time"`
}

// TableName 指定表名
func (UserProfileInterest) TableName() string {
	return "user_profile_interest"
}
