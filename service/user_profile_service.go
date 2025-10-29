package service

import (
	"AdvertRecommend/database"
	"AdvertRecommend/models"
	"errors"

	"gorm.io/gorm"
)

// UserProfileService 用户画像服务
type UserProfileService struct{}

// NewUserProfileService 创建用户画像服务实例
func NewUserProfileService() *UserProfileService {
	return &UserProfileService{}
}

// CreateUserProfile 创建用户画像
func (s *UserProfileService) CreateUserProfile(userID int64, gender, age int32, region, deviceType string) error {
	profile := &models.UserProfileBase{
		UserID:     userID,
		Gender:     gender,
		Age:        age,
		Region:     region,
		DeviceType: deviceType,
	}

	if err := database.DB.Create(profile).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUserProfile 更新用户画像
func (s *UserProfileService) UpdateUserProfile(userID int64, updates map[string]interface{}) error {
	result := database.DB.Model(&models.UserProfileBase{}).Where("user_id = ?", userID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user profile not found")
	}

	return nil
}

// GetUserProfile 获取用户画像
func (s *UserProfileService) GetUserProfile(userID int64) (*models.UserProfileBase, error) {
	var profile models.UserProfileBase
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user profile not found")
		}
		return nil, err
	}
	return &profile, nil
}

// DeleteUserProfile 删除用户画像
func (s *UserProfileService) DeleteUserProfile(userID int64) error {
	result := database.DB.Where("user_id = ?", userID).Delete(&models.UserProfileBase{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user profile not found")
	}

	return nil
}
