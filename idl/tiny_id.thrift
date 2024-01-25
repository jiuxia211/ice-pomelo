namespace go tiny_id

struct BaseResp {
    1: i64 code
    2: string msg
}

struct GetMaxIDRequest{
    1:i64 bizType
}

struct GetMaxIDResponse{
    1:BaseResp base
    2:i64 maxID
}

service TinyIDService{
    GetMaxIDResponse GetMaxID(1:GetMaxIDRequest req)
}