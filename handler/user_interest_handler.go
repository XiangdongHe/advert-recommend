package handler

import (
	"AdvertRecommend/kitex_gen/advert"
	"AdvertRecommend/models"
	"context"
	"log"
)

// ==================== 用户兴趣画像 CRUD ====================

// AddUserInterest 添加用户兴趣
func (s *AdvertServiceImpl) AddUserInterest(ctx context.Context, req *advert.AddUserInterestRequest) (*advert.AddUserInterestResponse, error) {
	log.Printf("AddUserInterest: %+v", req)

	id, err := s.userInterestService.AddUserInterest(req.UserId, req.Tag, req.Weight)
	if err != nil {
		return &advert.AddUserInterestResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.AddUserInterestResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Id: id,
	}, nil
}

// UpdateUserInterest 更新用户兴趣
func (s *AdvertServiceImpl) UpdateUserInterest(ctx context.Context, req *advert.UpdateUserInterestRequest) (*advert.UpdateUserInterestResponse, error) {
	log.Printf("UpdateUserInterest: %+v", req)

	err := s.userInterestService.UpdateUserInterest(req.Id, req.Weight)
	if err != nil {
		return &advert.UpdateUserInterestResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.UpdateUserInterestResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// GetUserInterests 获取用户兴趣列表
func (s *AdvertServiceImpl) GetUserInterests(ctx context.Context, req *advert.GetUserInterestsRequest) (*advert.GetUserInterestsResponse, error) {
	log.Printf("GetUserInterests: %+v", req)

	interests, err := s.userInterestService.GetUserInterests(req.UserId)
	if err != nil {
		return &advert.GetUserInterestsResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	userInterests := make([]*advert.UserInterest, 0, len(interests))
	for _, interest := range interests {
		userInterests = append(userInterests, convertUserInterest(interest))
	}

	return &advert.GetUserInterestsResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Interests: userInterests,
	}, nil
}

// DeleteUserInterest 删除用户兴趣
func (s *AdvertServiceImpl) DeleteUserInterest(ctx context.Context, req *advert.DeleteUserInterestRequest) (*advert.DeleteUserInterestResponse, error) {
	log.Printf("DeleteUserInterest: %+v", req)

	err := s.userInterestService.DeleteUserInterest(req.Id)
	if err != nil {
		return &advert.DeleteUserInterestResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.DeleteUserInterestResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// convertUserInterest 转换数据模型到 Thrift 模型
func convertUserInterest(interest *models.UserProfileInterest) *advert.UserInterest {
	return &advert.UserInterest{
		Id:         interest.ID,
		UserId:     interest.UserID,
		Tag:        interest.Tag,
		Weight:     interest.Weight,
		UpdateTime: interest.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}
