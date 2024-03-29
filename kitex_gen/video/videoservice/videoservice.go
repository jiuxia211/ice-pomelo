// Code generated by Kitex v0.8.0. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	video "github.com/jiuxia211/ice-pomelo/kitex_gen/video"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "videoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UploadVideo": kitex.NewMethodInfo(uploadVideoHandler, newVideoServiceUploadVideoArgs, newVideoServiceUploadVideoResult, false),
		"Feed":        kitex.NewMethodInfo(feedHandler, newVideoServiceFeedArgs, newVideoServiceFeedResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "video",
		"ServiceFilePath": `../../idl/video.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func uploadVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceUploadVideoArgs)
	realResult := result.(*video.VideoServiceUploadVideoResult)
	success, err := handler.(video.VideoService).UploadVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceUploadVideoArgs() interface{} {
	return video.NewVideoServiceUploadVideoArgs()
}

func newVideoServiceUploadVideoResult() interface{} {
	return video.NewVideoServiceUploadVideoResult()
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedArgs)
	realResult := result.(*video.VideoServiceFeedResult)
	success, err := handler.(video.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return video.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return video.NewVideoServiceFeedResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UploadVideo(ctx context.Context, req *video.UploadVideoRequest) (r *video.UploadVideoResponse, err error) {
	var _args video.VideoServiceUploadVideoArgs
	_args.Req = req
	var _result video.VideoServiceUploadVideoResult
	if err = p.c.Call(ctx, "UploadVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Feed(ctx context.Context, req *video.FeedRequest) (r *video.FeedResponse, err error) {
	var _args video.VideoServiceFeedArgs
	_args.Req = req
	var _result video.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
