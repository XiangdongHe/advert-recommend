package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gitee.com/HeXiangdong/AdvertRecommend/user-service/database"
	"gitee.com/HeXiangdong/AdvertRecommend/user-service/models"
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
func (s *UserProfileService) GetUserProfile(userId int64) (*models.UserProfileBase, error) {
	key := fmt.Sprintf("user:profile:%d", userId)
	var user models.UserProfileBase
	ctx := context.Background()
	// 尝试从 Redis 取
	val, err := database.RDB.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return &user, nil
		}
	}
	// 缓存未命中 → 访问数据库
	if err := database.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	// 写入缓存
	data, _ := json.Marshal(user)
	database.RDB.Set(ctx, key, data, time.Hour)
	return &user, nil
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
