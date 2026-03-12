package handler

import (
	"context"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/common"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/recommend"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/models"
)

// ==================== 广告创意 CRUD ====================

// CreateAdCreative 创建广告创意
func (s *RecommendServiceImpl) CreateAdCreative(ctx context.Context, req *recommend.CreateAdCreativeRequest) (*recommend.CreateAdCreativeResponse, error) {
	log.Printf("CreateAdCreative: %+v", req)

	creativeID, err := s.adCreativeService.CreateAdCreative(
		req.PlanId,
		req.CreativeType,
		req.MediaUrl,
		req.Title,
		req.Description,
	)

	if err != nil {
		return &recommend.CreateAdCreativeResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.CreateAdCreativeResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		CreativeId: creativeID,
	}, nil
}

// UpdateAdCreative 更新广告创意
func (s *RecommendServiceImpl) UpdateAdCreative(ctx context.Context, req *recommend.UpdateAdCreativeRequest) (*recommend.UpdateAdCreativeResponse, error) {
	log.Printf("UpdateAdCreative: %+v", req)

	updates := make(map[string]interface{})
	if req.PlanId != nil {
		updates["plan_id"] = *req.PlanId
	}
	if req.CreativeType != nil {
		updates["creative_type"] = *req.CreativeType
	}
	if req.MediaUrl != nil {
		updates["media_url"] = *req.MediaUrl
	}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	err := s.adCreativeService.UpdateAdCreative(req.CreativeId, updates)
	if err != nil {
		return &recommend.UpdateAdCreativeResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.UpdateAdCreativeResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// GetAdCreative 获取广告创意
func (s *RecommendServiceImpl) GetAdCreative(ctx context.Context, req *recommend.GetAdCreativeRequest) (*recommend.GetAdCreativeResponse, error) {
	log.Printf("GetAdCreative: %+v", req)

	creative, err := s.adCreativeService.GetAdCreative(req.CreativeId)
	if err != nil {
		return &recommend.GetAdCreativeResponse{
			BaseResp: &common.BaseResponse{
				Code:    404,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.GetAdCreativeResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		AdCreative: convertAdCreative(creative),
	}, nil
}

// ListAdCreatives 获取广告创意列表
func (s *RecommendServiceImpl) ListAdCreatives(ctx context.Context, req *recommend.ListAdCreativesRequest) (*recommend.ListAdCreativesResponse, error) {
	log.Printf("ListAdCreatives: %+v", req)

	creatives, total, err := s.adCreativeService.ListAdCreatives(int(req.Page), int(req.PageSize), req.PlanId)
	if err != nil {
		return &recommend.ListAdCreativesResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	adCreatives := make([]*recommend.AdCreative, 0, len(creatives))
	for _, creative := range creatives {
		adCreatives = append(adCreatives, convertAdCreative(creative))
	}

	return &recommend.ListAdCreativesResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		AdCreatives: adCreatives,
		Total:       total,
	}, nil
}

// DeleteAdCreative 删除广告创意
func (s *RecommendServiceImpl) DeleteAdCreative(ctx context.Context, req *recommend.DeleteAdCreativeRequest) (*recommend.DeleteAdCreativeResponse, error) {
	log.Printf("DeleteAdCreative: %+v", req)

	err := s.adCreativeService.DeleteAdCreative(req.CreativeId)
	if err != nil {
		return &recommend.DeleteAdCreativeResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.DeleteAdCreativeResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// convertAdCreative 转换数据模型到 Thrift 模型
func convertAdCreative(creative *models.AdCreative) *recommend.AdCreative {
	return &recommend.AdCreative{
		CreativeId:   creative.CreativeID,
		PlanId:       creative.PlanID,
		CreativeType: creative.CreativeType,
		MediaUrl:     creative.MediaURL,
		Title:        creative.Title,
		Description:  creative.Description,
		Status:       creative.Status,
		CreateTime:   creative.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:   creative.UpdateTime.Format("2006-01-02 15:04:05"),

		Weight: creative.Weight,
	}
}
