package handler

import (
	"AdvertRecommend/kitex_gen/advert"
	"AdvertRecommend/models"
	"context"
	"log"
)

// ==================== 用户行为日志 ====================

// CreateAdEvent 创建广告事件
func (s *AdvertServiceImpl) CreateAdEvent(ctx context.Context, req *advert.CreateAdEventRequest) (*advert.CreateAdEventResponse, error) {
	log.Printf("CreateAdEvent: %+v", req)

	eventID, err := s.adEventService.CreateAdEvent(
		req.UserId,
		req.CreativeId,
		req.EventType,
		req.Ts,
		req.Extra,
	)

	if err != nil {
		return &advert.CreateAdEventResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &advert.CreateAdEventResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		EventId: eventID,
	}, nil
}

// GetUserAdEvents 获取用户广告事件列表
func (s *AdvertServiceImpl) GetUserAdEvents(ctx context.Context, req *advert.GetUserAdEventsRequest) (*advert.GetUserAdEventsResponse, error) {
	log.Printf("GetUserAdEvents: %+v", req)

	events, total, err := s.adEventService.GetUserAdEvents(
		req.UserId,
		int(req.Page),
		int(req.PageSize),
		req.EventType,
	)

	if err != nil {
		return &advert.GetUserAdEventsResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	userAdEvents := make([]*advert.UserAdEvent, 0, len(events))
	for _, event := range events {
		userAdEvents = append(userAdEvents, convertUserAdEvent(event))
	}

	return &advert.GetUserAdEventsResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Events: userAdEvents,
		Total:  total,
	}, nil
}

// GetCreativeAdEvents 获取创意广告事件列表
func (s *AdvertServiceImpl) GetCreativeAdEvents(ctx context.Context, req *advert.GetCreativeAdEventsRequest) (*advert.GetCreativeAdEventsResponse, error) {
	log.Printf("GetCreativeAdEvents: %+v", req)

	events, total, err := s.adEventService.GetCreativeAdEvents(
		req.CreativeId,
		int(req.Page),
		int(req.PageSize),
		req.EventType,
	)

	if err != nil {
		return &advert.GetCreativeAdEventsResponse{
			BaseResp: &advert.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	creativeAdEvents := make([]*advert.UserAdEvent, 0, len(events))
	for _, event := range events {
		creativeAdEvents = append(creativeAdEvents, convertUserAdEvent(event))
	}

	return &advert.GetCreativeAdEventsResponse{
		BaseResp: &advert.BaseResponse{
			Code:    200,
			Message: "success",
		},
		Events: creativeAdEvents,
		Total:  total,
	}, nil
}

// convertUserAdEvent 转换数据模型到 Thrift 模型
func convertUserAdEvent(event *models.UserAdEventLog) *advert.UserAdEvent {
	return &advert.UserAdEvent{
		EventId:    event.EventID,
		UserId:     event.UserID,
		CreativeId: event.CreativeID,
		EventType:  event.EventType,
		Ts:         event.TS.Format("2006-01-02 15:04:05"),
		Extra:      event.Extra,
	}
}
