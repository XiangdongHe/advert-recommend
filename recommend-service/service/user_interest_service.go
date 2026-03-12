package service

import (
	"errors"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/database"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/models"
)

// UserInterestService 用户兴趣服务
type UserInterestService struct{}

// NewUserInterestService 创建用户兴趣服务实例
func NewUserInterestService() *UserInterestService {
	return &UserInterestService{}
}

// AddUserInterest 添加用户兴趣
func (s *UserInterestService) AddUserInterest(userID int64, tag string, weight float64) (int64, error) {
	interest := &models.UserProfileInterest{
		UserID: userID,
		Tag:    tag,
		Weight: weight,
	}

	if err := database.DB.Create(interest).Error; err != nil {
		return 0, err
	}

	return interest.ID, nil
}

// UpdateUserInterest 更新用户兴趣
func (s *UserInterestService) UpdateUserInterest(id int64, weight float64) error {
	result := database.DB.Model(&models.UserProfileInterest{}).Where("id = ?", id).Update("weight", weight)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user interest not found")
	}

	return nil
}

// GetUserInterests 获取用户兴趣列表
func (s *UserInterestService) GetUserInterests(userID int64) ([]*models.UserProfileInterest, error) {
	var interests []*models.UserProfileInterest
	if err := database.DB.Where("user_id = ?", userID).Find(&interests).Error; err != nil {
		return nil, err
	}
	return interests, nil
}

// DeleteUserInterest 删除用户兴趣
func (s *UserInterestService) DeleteUserInterest(id int64) error {
	result := database.DB.Where("id = ?", id).Delete(&models.UserProfileInterest{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user interest not found")
	}

	return nil
}
