namespace go video

struct BaseResp {
    1: i64 code
    2: string msg
}

struct Video{
    1: i64 id
    2: i64 uid
    2: string video_url
}
struct UploadVideoRequest{
    1: string token
    2: binary videoFile
}

struct UploadVideoResponse{
    1: BaseResp base
    2: i64 id
}

struct FeedRequest{
    1: string token
}

struct FeedResponse{
    1: BaseResp base
    2: list<Video> video_list
}