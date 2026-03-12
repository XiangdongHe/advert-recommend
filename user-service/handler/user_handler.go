package handler

import (
	"context"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/user-service/kitex_gen/common"
	"gitee.com/HeXiangdong/AdvertRecommend/user-service/kitex_gen/user"
	"gitee.com/HeXiangdong/AdvertRecommend/user-service/models"
	"gitee.com/HeXiangdong/AdvertRecommend/user-service/service"
)

// UserServiceImpl 实现 AdvertService 接口
type UserServiceImpl struct {
	userService *service.UserProfileService
}

// NewUserServiceImpl 创建服务实现实例
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		userService: service.NewUserProfileService(),
	}
}

// ==================== 用户画像 CRUD ====================

// CreateUserProfile 创建用户画像
func (s *UserServiceImpl) CreateUserProfile(ctx context.Context, req *user.CreateUserProfileRequest) (*user.CreateUserProfileResponse, error) {
	log.Printf("CreateUserProfile: %+v", req)

	err := s.userService.CreateUserProfile(
		req.UserId,
		req.Gender,
		req.Age,
		req.Region,
		req.DeviceType,
	)

	if err != nil {
		return &user.CreateUserProfileResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &user.CreateUserProfileResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// UpdateUserProfile 更新用户画像
func (s *UserServiceImpl) UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*user.UpdateUserProfileResponse, error) {
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

	err := s.userService.UpdateUserProfile(req.UserId, updates)
	if err != nil {
		return &user.UpdateUserProfileResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &user.UpdateUserProfileResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// GetUserProfile 获取用户画像
func (s *UserServiceImpl) GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error) {
	log.Printf("GetUserProfile: %+v", req)

	profile, err := s.userService.GetUserProfile(req.UserId)
	if err != nil {
		return &user.GetUserProfileResponse{
			BaseResp: &common.BaseResponse{
				Code:    404,
				Message: err.Error(),
			},
		}, nil
	}

	return &user.GetUserProfileResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		UserProfile: convertUserProfile(profile),
	}, nil
}

// DeleteUserProfile 删除用户画像
func (s *UserServiceImpl) DeleteUserProfile(ctx context.Context, req *user.DeleteUserProfileRequest) (*user.DeleteUserProfileResponse, error) {
	log.Printf("DeleteUserProfile: %+v", req)

	err := s.userService.DeleteUserProfile(req.UserId)
	if err != nil {
		return &user.DeleteUserProfileResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &user.DeleteUserProfileResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// convertUserProfile 转换数据模型到 Thrift 模型
func convertUserProfile(profile *models.UserProfileBase) *user.UserProfileBase {
	return &user.UserProfileBase{
		UserId:     profile.UserID,
		Gender:     profile.Gender,
		Age:        profile.Age,
		Region:     profile.Region,
		DeviceType: profile.DeviceType,
		CreateTime: profile.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: profile.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}
