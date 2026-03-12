package handler

import (
	"context"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/common"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/recommend"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/models"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/service"
)

// RecommendServiceImpl 实现 AdvertService 接口
type RecommendServiceImpl struct {
	adPlanService       *service.AdPlanService
	adCreativeService   *service.AdCreativeService
	userInterestService *service.UserInterestService
	adEventService      *service.UserAdEventService
}

// NewRecommendServiceImpl 创建服务实现实例
func NewRecommendServiceImpl() *RecommendServiceImpl {
	return &RecommendServiceImpl{
		adPlanService:       service.NewAdPlanService(),
		adCreativeService:   service.NewAdCreativeService(),
		userInterestService: service.NewUserInterestService(),
		adEventService:      service.NewUserAdEventService(),
	}
}

// ==================== 广告计划 CRUD ====================

// CreateAdPlan 创建广告计划
func (s *RecommendServiceImpl) CreateAdPlan(ctx context.Context, req *recommend.CreateAdPlanRequest) (*recommend.CreateAdPlanResponse, error) {
	log.Printf("CreateAdPlan: %+v", req)

	planID, err := s.adPlanService.CreateAdPlan(
		req.Name,
		req.Objective,
		req.Budget,
		req.BidPrice,
		req.TargetingRule,
		req.StartTime,
		req.EndTime,
	)

	if err != nil {
		return &recommend.CreateAdPlanResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.CreateAdPlanResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		PlanId: planID,
	}, nil
}

// UpdateAdPlan 更新广告计划
func (s *RecommendServiceImpl) UpdateAdPlan(ctx context.Context, req *recommend.UpdateAdPlanRequest) (*recommend.UpdateAdPlanResponse, error) {
	log.Printf("UpdateAdPlan: %+v", req)

	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Objective != nil {
		updates["objective"] = *req.Objective
	}
	if req.Budget != nil {
		updates["budget"] = *req.Budget
	}
	if req.BidPrice != nil {
		updates["bid_price"] = *req.BidPrice
	}
	if req.TargetingRule != nil {
		updates["targeting_rule"] = *req.TargetingRule
	}
	if req.StartTime != nil {
		updates["start_time"] = *req.StartTime
	}
	if req.EndTime != nil {
		updates["end_time"] = *req.EndTime
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	err := s.adPlanService.UpdateAdPlan(req.PlanId, updates)
	if err != nil {
		return &recommend.UpdateAdPlanResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.UpdateAdPlanResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// GetAdPlan 获取广告计划
func (s *RecommendServiceImpl) GetAdPlan(ctx context.Context, req *recommend.GetAdPlanRequest) (*recommend.GetAdPlanResponse, error) {
	log.Printf("GetAdPlan: %+v", req)

	plan, err := s.adPlanService.GetAdPlan(req.PlanId)
	if err != nil {
		return &recommend.GetAdPlanResponse{
			BaseResp: &common.BaseResponse{
				Code:    404,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.GetAdPlanResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		AdPlan: convertAdPlan(plan),
	}, nil
}

// ListAdPlans 获取广告计划列表
func (s *RecommendServiceImpl) ListAdPlans(ctx context.Context, req *recommend.ListAdPlansRequest) (*recommend.ListAdPlansResponse, error) {
	log.Printf("ListAdPlans: %+v", req)

	plans, total, err := s.adPlanService.ListAdPlans(int(req.Page), int(req.PageSize), req.Status)
	if err != nil {
		return &recommend.ListAdPlansResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	adPlans := make([]*recommend.AdPlan, 0, len(plans))
	for _, plan := range plans {
		adPlans = append(adPlans, convertAdPlan(plan))
	}

	return &recommend.ListAdPlansResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		AdPlans: adPlans,
		Total:   total,
	}, nil
}

// DeleteAdPlan 删除广告计划
func (s *RecommendServiceImpl) DeleteAdPlan(ctx context.Context, req *recommend.DeleteAdPlanRequest) (*recommend.DeleteAdPlanResponse, error) {
	log.Printf("DeleteAdPlan: %+v", req)

	err := s.adPlanService.DeleteAdPlan(req.PlanId)
	if err != nil {
		return &recommend.DeleteAdPlanResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &recommend.DeleteAdPlanResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
	}, nil
}

// convertAdPlan 转换数据模型到 Thrift 模型
func convertAdPlan(plan *models.AdPlan) *recommend.AdPlan {
	return &recommend.AdPlan{
		PlanId:        plan.PlanID,
		Name:          plan.Name,
		Objective:     plan.Objective,
		Budget:        plan.Budget,
		BidPrice:      plan.BidPrice,
		TargetingRule: plan.TargetingRule,
		StartTime:     plan.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:       plan.EndTime.Format("2006-01-02 15:04:05"),
		Status:        plan.Status,
		CreateTime:    plan.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:    plan.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}
