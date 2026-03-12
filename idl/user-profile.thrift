namespace go user

include "common.thrift"

// ==================== 用户画像相关 ====================
struct UserProfileBase {
    1: i64 userId
    2: i32 gender               // 0=unknown,1=male,2=female
    3: i32 age
    4: string region
    5: string deviceType
    6: string createTime
    7: string updateTime
}

struct GetUserProfileRequest {
    1: i64 userId
}

struct GetUserProfileResponse {
    1: common.BaseResponse baseResp
    2: UserProfileBase userProfile
}


service UserService {
    // 用户画像 CRUD
    GetUserProfileResponse GetUserProfile(1: GetUserProfileRequest req)
}