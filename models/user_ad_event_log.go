package models

import (
	"time"
)

// UserAdEventLog 用户广告行为日志
type UserAdEventLog struct {
	EventID    int64     `gorm:"primaryKey;autoIncrement;column:event_id" json:"event_id"`
	UserID     int64     `gorm:"not null;column:user_id;index:idx_user_ts,priority:1" json:"user_id"`
	CreativeID int64     `gorm:"not null;column:creative_id;index:idx_ad_ts,priority:1" json:"creative_id"`
	EventType  int32     `gorm:"type:tinyint;not null;column:event_type;comment:1=exposure,2=click,3=conversion" json:"event_type"`
	TS         time.Time `gorm:"type:datetime;not null;column:ts;index:idx_user_ts,priority:2;index:idx_ad_ts,priority:2" json:"ts"`
	Extra      string    `gorm:"type:json;column:extra" json:"extra"`
}

// TableName 指定表名
func (UserAdEventLog) TableName() string {
	return "user_ad_event_log"
}
