namespace go recommend

include "common.thrift"

// ==================== 广告计划相关 ====================
struct AdPlan {
    1: i64 planId
    2: string name
    3: string objective          // click/download/conversion
    4: double budget
    5: string bidPrice          // CPC/CPM/CPA 0.5元/0.01元/5元
    6: string targetingRule     // JSON string
    7: string startTime         // datetime string
    8: string endTime           // datetime string
    9: i32 status               // 1=active,0=paused,2=ended
    10: string createTime
    11: string updateTime
}

struct CreateAdPlanRequest {
    1: string name
    2: string objective
    3: double budget
    4: string bidPrice
    5: string targetingRule
    6: string startTime
    7: string endTime
}

struct CreateAdPlanResponse {
    1: common.BaseResponse baseResp
    2: i64 planId
}

struct UpdateAdPlanRequest {
    1: i64 planId
    2: optional string name
    3: optional string objective
    4: optional double budget
    5: optional string bidPrice
    6: optional string targetingRule
    7: optional string startTime
    8: optional string endTime
    9: optional i32 status
}

struct UpdateAdPlanResponse {
    1: common.BaseResponse baseResp
}

struct GetAdPlanRequest {
    1: i64 planId
}

struct GetAdPlanResponse {
    1: common.BaseResponse baseResp
    2: AdPlan adPlan
}

struct ListAdPlansRequest {
    1: i32 page
    2: i32 pageSize
    3: optional i32 status
}

struct ListAdPlansResponse {
    1: common.BaseResponse baseResp
    2: list<AdPlan> adPlans
    3: i64 total
}

struct DeleteAdPlanRequest {
    1: i64 planId
}

struct DeleteAdPlanResponse {
    1: common.BaseResponse baseResp
}

// ==================== 广告创意相关 ====================
struct AdCreative {
    1: i64 creativeId
    2: i64 planId
    3: i32 creativeType         // 1=image,2=video,3=text
    4: string mediaUrl
    5: string title
    6: string description
    7: i32 status
    8: string createTime
    9: string updateTime
    10:double weight
}

struct CreateAdCreativeRequest {
    1: i64 planId
    2: i32 creativeType
    3: string mediaUrl
    4: string title
    5: string description
}

struct CreateAdCreativeResponse {
    1: common.BaseResponse baseResp
    2: i64 creativeId
}

struct UpdateAdCreativeRequest {
    1: i64 creativeId
    2: optional i64 planId
    3: optional i32 creativeType
    4: optional string mediaUrl
    5: optional string title
    6: optional string description
    7: optional i32 status
}

struct UpdateAdCreativeResponse {
    1: common.BaseResponse baseResp
}

struct GetAdCreativeRequest {
    1: i64 creativeId
}

struct GetAdCreativeResponse {
    1: common.BaseResponse baseResp
    2: AdCreative adCreative
}

struct ListAdCreativesRequest {
    1: i32 page
    2: i32 pageSize
    3: optional i64 planId
}

struct ListAdCreativesResponse {
    1: common.BaseResponse baseResp
    2: list<AdCreative> adCreatives
    3: i64 total
}

struct DeleteAdCreativeRequest {
    1: i64 creativeId
}

struct DeleteAdCreativeResponse {
    1: common.BaseResponse baseResp
}

// ==================== 用户兴趣画像相关 ====================
struct UserInterest {
    1: i64 id
    2: i64 userId
    3: string tag
    4: double weight
    5: string updateTime
}

struct AddUserInterestRequest {
    1: i64 userId
    2: string tag
    3: double weight
}

struct AddUserInterestResponse {
    1: common.BaseResponse baseResp
    2: i64 id
}

struct UpdateUserInterestRequest {
    1: i64 id
    2: double weight
}

struct UpdateUserInterestResponse {
    1: common.BaseResponse baseResp
}

struct GetUserInterestsRequest {
    1: i64 userId
}

struct GetUserInterestsResponse {
    1: common.BaseResponse baseResp
    2: list<UserInterest> interests
}

struct DeleteUserInterestRequest {
    1: i64 id
}

struct DeleteUserInterestResponse {
    1: common.BaseResponse baseResp
}

// ==================== 用户行为日志相关 ====================
struct UserAdEvent {
    1: i64 eventId
    2: i64 userId
    3: i64 creativeId
    4: i32 eventType            // 1=exposure,2=click,3=conversion
    5: string ts
    6: string extra             // JSON string
}

struct CreateAdEventRequest {
    1: i64 userId
    2: i64 creativeId
    3: i32 eventType
    4: string ts
    5: string extra
}

struct CreateAdEventResponse {
    1: common.BaseResponse baseResp
    2: i64 eventId
}

struct GetUserAdEventsRequest {
    1: i64 userId
    2: i32 page
    3: i32 pageSize
    4: optional i32 eventType
}

struct GetUserAdEventsResponse {
    1: common.BaseResponse baseResp
    2: list<UserAdEvent> events
    3: i64 total
}

struct GetCreativeAdEventsRequest {
    1: i64 creativeId
    2: i32 page
    3: i32 pageSize
    4: optional i32 eventType
}

struct GetCreativeAdEventsResponse {
    1: common.BaseResponse baseResp
    2: list<UserAdEvent> events
    3: i64 total
}

// 广告查询相关
struct GetAdvertRecommendRequest {
    1: i64 userId
}

struct GetAdvertRecommendResponse {
    1: common.BaseResponse baseResp
    2: list<AdCreative> adverts
    3: i64 total
}

// ==================== 服务定义 ====================
service RecommendService {
    // 广告计划 CRUD
    CreateAdPlanResponse CreateAdPlan(1: CreateAdPlanRequest req)
    UpdateAdPlanResponse UpdateAdPlan(1: UpdateAdPlanRequest req)
    GetAdPlanResponse GetAdPlan(1: GetAdPlanRequest req)
    ListAdPlansResponse ListAdPlans(1: ListAdPlansRequest req)
    DeleteAdPlanResponse DeleteAdPlan(1: DeleteAdPlanRequest req)
    
    // 广告创意 CRUD
    CreateAdCreativeResponse CreateAdCreative(1: CreateAdCreativeRequest req)
    UpdateAdCreativeResponse UpdateAdCreative(1: UpdateAdCreativeRequest req)
    GetAdCreativeResponse GetAdCreative(1: GetAdCreativeRequest req)
    ListAdCreativesResponse ListAdCreatives(1: ListAdCreativesRequest req)
    DeleteAdCreativeResponse DeleteAdCreative(1: DeleteAdCreativeRequest req)
    
    // 用户兴趣画像 CRUD
    AddUserInterestResponse AddUserInterest(1: AddUserInterestRequest req)
    UpdateUserInterestResponse UpdateUserInterest(1: UpdateUserInterestRequest req)
    GetUserInterestsResponse GetUserInterests(1: GetUserInterestsRequest req)
    DeleteUserInterestResponse DeleteUserInterest(1: DeleteUserInterestRequest req)
    
    // 用户行为日志 暂时用不到
    CreateAdEventResponse CreateAdEvent(1: CreateAdEventRequest req)
    GetUserAdEventsResponse GetUserAdEvents(1: GetUserAdEventsRequest req)
    GetCreativeAdEventsResponse GetCreativeAdEvents(1: GetCreativeAdEventsRequest req)

    // 获取用户的广告
    GetAdvertRecommendResponse GetAdvertRecommend(1: GetAdvertRecommendRequest req)
}
