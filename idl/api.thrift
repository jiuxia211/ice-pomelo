namespace go api

struct BaseResp {
    1: i64 code
    2: string msg
}

struct User {
    1: i64 id
    2: string name
    3: string avatar_url
}

struct RegisterRequest{
    1: string username
    2: string password
    3: string email
    4: string verificationCode
}

struct RegisterResponse{
    1: BaseResp base
    2: i64 id
}

struct LoginRequest{
    1: string username
    2: string password
}

struct LoginResponse{
    1: BaseResp base
    2: i64 id
    3: string token
}

struct SendVerificationCodeRequest{
    1: string email
}

struct SendVerificationCodeResponse{
    1: BaseResp base
}

struct GetUserInfoRequest{
    1: string token(api.header="Authorization")
    2: i64 id
}

struct GetUserInfoResponse{
    1: BaseResp base
    2: User user
}

struct UploadUserAvatarRequest{
    1: string token(api.header="Authorization")
}

struct UploadUserAvatarResponse{
    1: BaseResp base
    2: User user
}

struct Video{
    1: i64 id
    2: i64 uid
    3: string video_url
    4: string cover_url
    5: string title
    6: string introduction
}

struct UploadVideoRequest{
    1: string token(api.header="Authorization")
    2: binary videoFile
    3: string video_format
    4: binary coverFile
    5: string cover_format
    6: string title
    7: string introduction
   
}

struct UploadVideoResponse{
    1: BaseResp base
    2: i64 id
}

struct FeedRequest{
    1: string token(api.header="Authorization")
}

struct FeedResponse{
    1: BaseResp base
    2: list<Video> video_list
}


service UserService{
    RegisterResponse Register(1:RegisterRequest req) (api.post="/pomelo/user/register")
    LoginResponse Login(1:LoginRequest req) (api.post="/pomelo/user/login")
    SendVerificationCodeResponse SendVerificationCode(1:SendVerificationCodeRequest req) (api.post="/pomelo/user/verification-code")
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req) (api.get="/pomelo/user")
    UploadUserAvatarResponse UploadUserAvatar(1:UploadUserAvatarRequest req)(api.put="/pomelo/user/avatar")
}

service VideoService{
    UploadVideoResponse UploadVideo(1:UploadVideoRequest req)(api.put="/pomelo/video")
    FeedResponse Feed(1:FeedRequest req)(api.get="/pomelo/video/feed")
}
