package models

import (
	"time"
)

// UserProfileBase 用户静态画像（基础属性）
type UserProfileBase struct {
	UserID     int64     `gorm:"primaryKey;column:user_id" json:"user_id"`
	Gender     int32     `gorm:"type:tinyint;column:gender;comment:0=unknown,1=male,2=female" json:"gender"`
	Age        int32     `gorm:"column:age" json:"age"`
	Region     string    `gorm:"type:varchar(64);column:region" json:"region"`
	DeviceType string    `gorm:"type:varchar(64);column:device_type" json:"device_type"`
	CreateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_time" json:"update_time"`
}

// TableName 指定表名
func (UserProfileBase) TableName() string {
	return "user_profile_base"
}
