package service

import (
	"AdvertRecommend/database"
	"AdvertRecommend/models"
	"errors"
	"time"
)

// UserAdEventService 用户广告事件服务
type UserAdEventService struct{}

// NewUserAdEventService 创建用户广告事件服务实例
func NewUserAdEventService() *UserAdEventService {
	return &UserAdEventService{}
}

// CreateAdEvent 创建广告事件
func (s *UserAdEventService) CreateAdEvent(userID, creativeID int64, eventType int32, ts, extra string) (int64, error) {
	tsObj, err := time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		return 0, errors.New("invalid timestamp format")
	}

	event := &models.UserAdEventLog{
		UserID:     userID,
		CreativeID: creativeID,
		EventType:  eventType,
		TS:         tsObj,
		Extra:      extra,
	}

	if err := database.DB.Create(event).Error; err != nil {
		return 0, err
	}

	return event.EventID, nil
}

// GetUserAdEvents 获取用户广告事件列表
func (s *UserAdEventService) GetUserAdEvents(userID int64, page, pageSize int, eventType *int32) ([]*models.UserAdEventLog, int64, error) {
	var events []*models.UserAdEventLog
	var total int64

	query := database.DB.Model(&models.UserAdEventLog{}).Where("user_id = ?", userID)
	if eventType != nil {
		query = query.Where("event_type = ?", *eventType)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Order("ts DESC").Offset(offset).Limit(pageSize).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

// GetCreativeAdEvents 获取创意广告事件列表
func (s *UserAdEventService) GetCreativeAdEvents(creativeID int64, page, pageSize int, eventType *int32) ([]*models.UserAdEventLog, int64, error) {
	var events []*models.UserAdEventLog
	var total int64

	query := database.DB.Model(&models.UserAdEventLog{}).Where("creative_id = ?", creativeID)
	if eventType != nil {
		query = query.Where("event_type = ?", *eventType)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Order("ts DESC").Offset(offset).Limit(pageSize).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}
