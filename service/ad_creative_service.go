package service

import (
	"AdvertRecommend/database"
	"AdvertRecommend/models"
	"errors"

	"gorm.io/gorm"
)

// AdCreativeService 广告创意服务
type AdCreativeService struct{}

// NewAdCreativeService 创建广告创意服务实例
func NewAdCreativeService() *AdCreativeService {
	return &AdCreativeService{}
}

// CreateAdCreative 创建广告创意
func (s *AdCreativeService) CreateAdCreative(planID int64, creativeType int32, mediaURL, title, description string) (int64, error) {
	creative := &models.AdCreative{
		PlanID:       planID,
		CreativeType: creativeType,
		MediaURL:     mediaURL,
		Title:        title,
		Description:  description,
		Status:       1, // 默认激活
	}

	if err := database.DB.Create(creative).Error; err != nil {
		return 0, err
	}

	return creative.CreativeID, nil
}

// UpdateAdCreative 更新广告创意
func (s *AdCreativeService) UpdateAdCreative(creativeID int64, updates map[string]interface{}) error {
	result := database.DB.Model(&models.AdCreative{}).Where("creative_id = ?", creativeID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("creative not found")
	}

	return nil
}

// GetAdCreative 获取广告创意
func (s *AdCreativeService) GetAdCreative(creativeID int64) (*models.AdCreative, error) {
	var creative models.AdCreative
	if err := database.DB.Where("creative_id = ?", creativeID).First(&creative).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("creative not found")
		}
		return nil, err
	}
	return &creative, nil
}

// ListAdCreatives 获取广告创意列表
func (s *AdCreativeService) ListAdCreatives(page, pageSize int, planID *int64) ([]*models.AdCreative, int64, error) {
	var creatives []*models.AdCreative
	var total int64

	query := database.DB.Model(&models.AdCreative{})
	if planID != nil {
		query = query.Where("plan_id = ?", *planID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&creatives).Error; err != nil {
		return nil, 0, err
	}

	return creatives, total, nil
}

// DeleteAdCreative 删除广告创意
func (s *AdCreativeService) DeleteAdCreative(creativeID int64) error {
	result := database.DB.Where("creative_id = ?", creativeID).Delete(&models.AdCreative{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("creative not found")
	}

	return nil
}

// 获取推荐广告列表
func (s *AdCreativeService) GetAdvertRecommend(userId int64) ([]*models.AdCreative, int64, error) {
	var creatives []*models.AdCreative
	var total int64
	// TODO 完成广告推荐的逻辑

	// 1.基于规则匹配到的广告集合

	// 2.基于内容匹配到的广告集合

	// 3.基于协同过滤匹配到的广告集合

	// 4.基于向量召回匹配到的广告集合

	return creatives, total, nil
}
