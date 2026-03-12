package service

import (
	"errors"
	"time"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/database"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/models"
	"gorm.io/gorm"
)

// AdPlanService 广告计划服务
type AdPlanService struct{}

// NewAdPlanService 创建广告计划服务实例
func NewAdPlanService() *AdPlanService {
	return &AdPlanService{}
}

// CreateAdPlan 创建广告计划
func (s *AdPlanService) CreateAdPlan(name, objective string, budget float64, bidPrice, targetingRule, startTime, endTime string) (int64, error) {
	startTimeObj, err := time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		return 0, errors.New("invalid start time format")
	}

	endTimeObj, err := time.Parse("2006-01-02 15:04:05", endTime)
	if err != nil {
		return 0, errors.New("invalid end time format")
	}

	plan := &models.AdPlan{
		Name:          name,
		Objective:     objective,
		Budget:        budget,
		BidPrice:      bidPrice,
		TargetingRule: targetingRule,
		StartTime:     startTimeObj,
		EndTime:       endTimeObj,
		Status:        1, // 默认激活
	}

	if err := database.DB.Create(plan).Error; err != nil {
		return 0, err
	}

	return plan.PlanID, nil
}

// UpdateAdPlan 更新广告计划
func (s *AdPlanService) UpdateAdPlan(planID int64, updates map[string]interface{}) error {
	// 处理时间字段
	if startTime, ok := updates["start_time"].(string); ok {
		startTimeObj, err := time.Parse("2006-01-02 15:04:05", startTime)
		if err != nil {
			return errors.New("invalid start time format")
		}
		updates["start_time"] = startTimeObj
	}

	if endTime, ok := updates["end_time"].(string); ok {
		endTimeObj, err := time.Parse("2006-01-02 15:04:05", endTime)
		if err != nil {
			return errors.New("invalid end time format")
		}
		updates["end_time"] = endTimeObj
	}

	result := database.DB.Model(&models.AdPlan{}).Where("plan_id = ?", planID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("plan not found")
	}

	return nil
}

// GetAdPlan 获取广告计划
func (s *AdPlanService) GetAdPlan(planID int64) (*models.AdPlan, error) {
	var plan models.AdPlan
	if err := database.DB.Where("plan_id = ?", planID).First(&plan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("plan not found")
		}
		return nil, err
	}
	return &plan, nil
}

// ListAdPlans 获取广告计划列表
func (s *AdPlanService) ListAdPlans(page, pageSize int, status *int32) ([]*models.AdPlan, int64, error) {
	var plans []*models.AdPlan
	var total int64

	query := database.DB.Model(&models.AdPlan{})
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&plans).Error; err != nil {
		return nil, 0, err
	}

	return plans, total, nil
}

// DeleteAdPlan 删除广告计划
func (s *AdPlanService) DeleteAdPlan(planID int64) error {
	result := database.DB.Where("plan_id = ?", planID).Delete(&models.AdPlan{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("plan not found")
	}

	return nil
}
