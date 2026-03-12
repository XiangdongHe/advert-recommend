package handler

import (
	"context"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/advert"
)

// GetAdCreative 获取广告创意
func (s *AdvertServiceImpl) GetAdvertRecommend(ctx context.Context, req *advert.GetAdvertRecommendRequest) (*advert.GetAdvertRecommendResponse, error) {
	log.Printf("GetAdvertRecommend: %+v", req)

	creatives, total, err := s.adCreativeService.GetAdvertRecommend(req.GetUserId())
	if err != nil {
		return &advert.GetAdvertRecommendResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	adCreatives := make([]*advert.AdCreative, 0, len(creatives))
	for _, ad := range creatives {
		adCreatives = append(adCreatives, convertAdCreative(ad))
	}

	return &advert.GetAdvertRecommendResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Adverts: adCreatives,
		Total:   total,
	}, nil
}
