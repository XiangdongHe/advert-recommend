package handler

import (
	"context"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/common"
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/recommend"
)

// GetAdCreative 获取广告创意
func (s *RecommendServiceImpl) GetAdvertRecommend(ctx context.Context, req *recommend.GetAdvertRecommendRequest) (*recommend.GetAdvertRecommendResponse, error) {
	log.Printf("GetAdvertRecommend: %+v", req)

	creatives, total, err := s.adCreativeService.GetAdvertRecommend(req.GetUserId())
	if err != nil {
		return &recommend.GetAdvertRecommendResponse{
			BaseResp: &common.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	adCreatives := make([]*recommend.AdCreative, 0, len(creatives))
	for _, ad := range creatives {
		adCreatives = append(adCreatives, convertAdCreative(ad))
	}

	return &recommend.GetAdvertRecommendResponse{
		BaseResp: &common.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Adverts: adCreatives,
		Total:   total,
	}, nil
}
