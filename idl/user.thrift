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

struct CreateUserProfileRequest {
    1: i64 userId
    2: i32 gender
    3: i32 age
    4: string region
    5: string deviceType
}

struct CreateUserProfileResponse {
    1: common.BaseResponse baseResp
}

struct UpdateUserProfileRequest {
    1: i64 userId
    2: optional i32 gender
    3: optional i32 age
    4: optional string region
    5: optional string deviceType
}

struct UpdateUserProfileResponse {
    1: common.BaseResponse baseResp
}

struct GetUserProfileRequest {
    1: i64 userId
}

struct GetUserProfileResponse {
    1: common.BaseResponse baseResp
    2: UserProfileBase userProfile
}

struct DeleteUserProfileRequest {
    1: i64 userId
}

struct DeleteUserProfileResponse {
    1: common.BaseResponse baseResp
}

service UserService {
    // 用户画像 CRUD
    CreateUserProfileResponse CreateUserProfile(1: CreateUserProfileRequest req)
    UpdateUserProfileResponse UpdateUserProfile(1: UpdateUserProfileRequest req)
    GetUserProfileResponse GetUserProfile(1: GetUserProfileRequest req)
    DeleteUserProfileResponse DeleteUserProfile(1: DeleteUserProfileRequest req)
}