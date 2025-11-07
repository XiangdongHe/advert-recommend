package service

import (
	"AdvertRecommend/database"
	"AdvertRecommend/models"
	"encoding/json"
	"errors"
	"sort"

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
	var ansCreatives []*models.AdCreative
	var ansAdPlans []*models.AdPlan
	var total int64
	// TODO 完成广告推荐的逻辑
	// 提取用户的基本信息
	var user models.UserProfileBase
	if err := database.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New("user not found")
		}
		return nil, 0, err
	}
	var userInterests []*models.UserProfileInterest
	if err := database.DB.Where("user_id = ?", userId).Find(&userInterests).Error; err != nil {
		return nil, 0, err
	}
	// 提取用户兴趣标签
	var tags []string
	for _, ui := range userInterests {
		tags = append(tags, ui.Tag)
	}

	// 1.基于规则匹配到的广告集合
	interestAdPlans, _ := GetInterestAdPlans(userInterests)
	ansAdPlans = append(ansAdPlans, interestAdPlans...)
	// 2.基于内容匹配到的广告集合

	// 3.基于协同过滤匹配到的广告集合

	// 4.基于向量召回匹配到的广告集合

	// 根据兴趣权重进行粗排
	interestWeight := make(map[string]float64)
	for _, ui := range userInterests {
		interestWeight[ui.Tag] = ui.Weight
	}

	for _, adPlan := range ansAdPlans {
		var targeting map[string]string
		if err := json.Unmarshal([]byte(adPlan.TargetingRule), &targeting); err != nil {
			continue
		}
		interest, ok := targeting["interest"]
		if !ok {
			continue
		}
		w := interestWeight[interest]
		for _, ad := range adPlan.Creatives {
			ad.Weight = w
			ansCreatives = append(ansCreatives, ad)
		}
	}

	sort.SliceStable(ansCreatives, func(i, j int) bool {
		return ansCreatives[i].Weight > ansCreatives[j].Weight
	})

	// TODO 根据点击、浏览、点赞模型预测，进行粗排

	return ansCreatives, total, nil
}

func GetInterestAdPlans(userInterests []*models.UserProfileInterest) ([]*models.AdPlan, error) {
	var adPlans []*models.AdPlan

	// 查出符合兴趣的广告计划
	query := database.DB.Model(&models.AdPlan{})
	for i, it := range userInterests {
		if i == 0 {
			query = query.Where("targeting_rule LIKE ?", "%"+it.Tag+"%")
		} else {
			query = query.Or("targeting_rule LIKE ?", "%"+it.Tag+"%")
		}
	}
	if err := query.Where("status = ?", 1).Find(&adPlans).Error; err != nil {
		return nil, err
	}
	if len(adPlans) == 0 {
		return adPlans, nil
	}

	// 查出所有相关广告创意
	var planIDs []int64
	for _, p := range adPlans {
		planIDs = append(planIDs, p.PlanID)
	}
	var adCreatives []*models.AdCreative
	if err := database.DB.Where("plan_id IN ?", planIDs).
		Where("status = ?", 1).Find(&adCreatives).Error; err != nil {
		return nil, err
	}

	// 将广告创意分配给对应的计划
	creativeMap := make(map[int64][]*models.AdCreative)
	for _, c := range adCreatives {
		creativeMap[c.PlanID] = append(creativeMap[c.PlanID], c)
	}
	for _, plan := range adPlans {
		plan.Creatives = creativeMap[plan.PlanID]
	}

	return adPlans, nil
}
