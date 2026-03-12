package handler

import (
	"context"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/common"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/recommend"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/models"
)

// ==================== 用户兴趣画像 CRUD ====================

// AddUserInterest 添加用户兴趣
func (s *RecommendServiceImpl) AddUserInterest(ctx context.Context, req *recommend.AddUserInterestRequest) (*recommend.AddUserInterestResponse, error) {
	log.Printf("AddUserInterest: %+v", req)

	id, err := s.userInterestService.AddUserInterest(req.UserId, req.Tag, req.Weight)
	if err != nil {
		return &recommend.AddUserInterestResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.AddUserInterestResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Id: id,
	}, nil
}

// UpdateUserInterest 更新用户兴趣
func (s *RecommendServiceImpl) UpdateUserInterest(ctx context.Context, req *recommend.UpdateUserInterestRequest) (*recommend.UpdateUserInterestResponse, error) {
	log.Printf("UpdateUserInterest: %+v", req)

	err := s.userInterestService.UpdateUserInterest(req.Id, req.Weight)
	if err != nil {
		return &recommend.UpdateUserInterestResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.UpdateUserInterestResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// GetUserInterests 获取用户兴趣列表
func (s *RecommendServiceImpl) GetUserInterests(ctx context.Context, req *recommend.GetUserInterestsRequest) (*recommend.GetUserInterestsResponse, error) {
	log.Printf("GetUserInterests: %+v", req)

	interests, err := s.userInterestService.GetUserInterests(req.UserId)
	if err != nil {
		return &recommend.GetUserInterestsResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	userInterests := make([]*recommend.UserInterest, 0, len(interests))
	for _, interest := range interests {
		userInterests = append(userInterests, convertUserInterest(interest))
	}

	return &recommend.GetUserInterestsResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Interests: userInterests,
	}, nil
}

// DeleteUserInterest 删除用户兴趣
func (s *RecommendServiceImpl) DeleteUserInterest(ctx context.Context, req *recommend.DeleteUserInterestRequest) (*recommend.DeleteUserInterestResponse, error) {
	log.Printf("DeleteUserInterest: %+v", req)

	err := s.userInterestService.DeleteUserInterest(req.Id)
	if err != nil {
		return &recommend.DeleteUserInterestResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.DeleteUserInterestResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// convertUserInterest 转换数据模型到 Thrift 模型
func convertUserInterest(interest *models.UserProfileInterest) *recommend.UserInterest {
	return &recommend.UserInterest{
		Id:         interest.ID,
		UserId:     interest.UserID,
		Tag:        interest.Tag,
		Weight:     interest.Weight,
		UpdateTime: interest.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}
