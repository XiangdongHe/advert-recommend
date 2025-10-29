package handler

import (
	"AdvertRecommend/kitex_gen/advert"
	"AdvertRecommend/models"
	"context"
	"log"
)

// ==================== 用户画像 CRUD ====================

// CreateUserProfile 创建用户画像
func (s *AdvertServiceImpl) CreateUserProfile(ctx context.Context, req *advert.CreateUserProfileRequest) (*advert.CreateUserProfileResponse, error) {
	log.Printf("CreateUserProfile: %+v", req)

	err := s.userProfileService.CreateUserProfile(
		req.UserId,
		req.Gender,
		req.Age,
		req.Region,
		req.DeviceType,
	)

	if err != nil {
		return &advert.CreateUserProfileResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.CreateUserProfileResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// UpdateUserProfile 更新用户画像
func (s *AdvertServiceImpl) UpdateUserProfile(ctx context.Context, req *advert.UpdateUserProfileRequest) (*advert.UpdateUserProfileResponse, error) {
	log.Printf("UpdateUserProfile: %+v", req)

	updates := make(map[string]interface{})
	if req.Gender != nil {
		updates["gender"] = *req.Gender
	}
	if req.Age != nil {
		updates["age"] = *req.Age
	}
	if req.Region != nil {
		updates["region"] = *req.Region
	}
	if req.DeviceType != nil {
		updates["device_type"] = *req.DeviceType
	}

	err := s.userProfileService.UpdateUserProfile(req.UserId, updates)
	if err != nil {
		return &advert.UpdateUserProfileResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.UpdateUserProfileResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// GetUserProfile 获取用户画像
func (s *AdvertServiceImpl) GetUserProfile(ctx context.Context, req *advert.GetUserProfileRequest) (*advert.GetUserProfileResponse, error) {
	log.Printf("GetUserProfile: %+v", req)

	profile, err := s.userProfileService.GetUserProfile(req.UserId)
	if err != nil {
		return &advert.GetUserProfileResponse{
			BaseResp: &advert.BaseResponse{
				Code:    404,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.GetUserProfileResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		UserProfile: convertUserProfile(profile),
	}, nil
}

// DeleteUserProfile 删除用户画像
func (s *AdvertServiceImpl) DeleteUserProfile(ctx context.Context, req *advert.DeleteUserProfileRequest) (*advert.DeleteUserProfileResponse, error) {
	log.Printf("DeleteUserProfile: %+v", req)

	err := s.userProfileService.DeleteUserProfile(req.UserId)
	if err != nil {
		return &advert.DeleteUserProfileResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.DeleteUserProfileResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// convertUserProfile 转换数据模型到 Thrift 模型
func convertUserProfile(profile *models.UserProfileBase) *advert.UserProfileBase {
	return &advert.UserProfileBase{
		UserId:     profile.UserID,
		Gender:     profile.Gender,
		Age:        profile.Age,
		Region:     profile.Region,
		DeviceType: profile.DeviceType,
		CreateTime: profile.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: profile.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}
