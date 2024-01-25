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

service UserService{
    RegisterResponse Register(1:RegisterRequest req) (api.post="/pomelo/user/register")
    LoginResponse Login(1:LoginRequest req) (api.post="/pomelo/user/login")
    SendVerificationCodeResponse SendVerificationCode(1:SendVerificationCodeRequest req) (api.post="/pomelo/user/verification-code")
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req) (api.get="/pomelo/user")
    UploadUserAvatarResponse UploadUserAvatar(1:UploadUserAvatarRequest req)(api.put="/pomelo/user/avatar")
}