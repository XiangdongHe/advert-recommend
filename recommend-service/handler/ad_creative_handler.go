package handler

import (
	"context"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/advert"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/models"
)

// ==================== 广告创意 CRUD ====================

// CreateAdCreative 创建广告创意
func (s *AdvertServiceImpl) CreateAdCreative(ctx context.Context, req *advert.CreateAdCreativeRequest) (*advert.CreateAdCreativeResponse, error) {
	log.Printf("CreateAdCreative: %+v", req)

	creativeID, err := s.adCreativeService.CreateAdCreative(
		req.PlanId,
		req.CreativeType,
		req.MediaUrl,
		req.Title,
		req.Description,
	)

	if err != nil {
		return &advert.CreateAdCreativeResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.CreateAdCreativeResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		CreativeId: creativeID,
	}, nil
}

// UpdateAdCreative 更新广告创意
func (s *AdvertServiceImpl) UpdateAdCreative(ctx context.Context, req *advert.UpdateAdCreativeRequest) (*advert.UpdateAdCreativeResponse, error) {
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
		return &advert.UpdateAdCreativeResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.UpdateAdCreativeResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// GetAdCreative 获取广告创意
func (s *AdvertServiceImpl) GetAdCreative(ctx context.Context, req *advert.GetAdCreativeRequest) (*advert.GetAdCreativeResponse, error) {
	log.Printf("GetAdCreative: %+v", req)

	creative, err := s.adCreativeService.GetAdCreative(req.CreativeId)
	if err != nil {
		return &advert.GetAdCreativeResponse{
			BaseResp: &advert.BaseResponse{
				Code:    404,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.GetAdCreativeResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		AdCreative: convertAdCreative(creative),
	}, nil
}

// ListAdCreatives 获取广告创意列表
func (s *AdvertServiceImpl) ListAdCreatives(ctx context.Context, req *advert.ListAdCreativesRequest) (*advert.ListAdCreativesResponse, error) {
	log.Printf("ListAdCreatives: %+v", req)

	creatives, total, err := s.adCreativeService.ListAdCreatives(int(req.Page), int(req.PageSize), req.PlanId)
	if err != nil {
		return &advert.ListAdCreativesResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	adCreatives := make([]*advert.AdCreative, 0, len(creatives))
	for _, creative := range creatives {
		adCreatives = append(adCreatives, convertAdCreative(creative))
	}

	return &advert.ListAdCreativesResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		AdCreatives: adCreatives,
		Total:       total,
	}, nil
}

// DeleteAdCreative 删除广告创意
func (s *AdvertServiceImpl) DeleteAdCreative(ctx context.Context, req *advert.DeleteAdCreativeRequest) (*advert.DeleteAdCreativeResponse, error) {
	log.Printf("DeleteAdCreative: %+v", req)

	err := s.adCreativeService.DeleteAdCreative(req.CreativeId)
	if err != nil {
		return &advert.DeleteAdCreativeResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.DeleteAdCreativeResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// convertAdCreative 转换数据模型到 Thrift 模型
func convertAdCreative(creative *models.AdCreative) *advert.AdCreative {
	return &advert.AdCreative{
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
