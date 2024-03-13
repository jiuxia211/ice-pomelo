package rpc

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video/videoservice"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

func InitVideoRPC() {
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithHostPorts("127.0.0.1:8212"))
	if err != nil {
		klog.Fatalf("video init rpc error: %v", err)
	}
	videoClient = c
}

func VideoUpload(ctx context.Context, req *video.UploadVideoRequest) (int64, error) {
	resp, err := videoClient.UploadVideo(ctx, req)

	if err != nil {
		return -1, err
	}

	if resp.Base.Code != errz.SuccessCode {
		return 0, errors.New(resp.Base.Msg)
	}

	return resp.Id, nil
}

func VideoFeed(ctx context.Context, req *video.FeedRequest) ([]*video.Video, error) {
	resp, err := videoClient.Feed(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errz.SuccessCode {
		return nil, errors.New(resp.Base.Msg)
	}

	return resp.VideoList, nil
}
