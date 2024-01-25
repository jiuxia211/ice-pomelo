namespace go user

struct BaseResp {
    1: i64 code
    2: string msg
}

struct User {
    1: i64 id
    2: string name
    3: string avatar
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
    1: string token
    2: i64 id
}

struct GetUserInfoResponse{
    1: BaseResp base
    2: User user
}

service UserService{
    RegisterResponse Register(1:RegisterRequest req)
    LoginResponse Login(1:LoginRequest req)
    SendVerificationCodeResponse SendVerificationCode(1:SendVerificationCodeRequest req)
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req)
}