package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jiuxia211/ice-pomelo/cmd/video/dal/db"
	"github.com/jiuxia211/ice-pomelo/cmd/video/pack"
	"github.com/jiuxia211/ice-pomelo/cmd/video/service"
	"github.com/jiuxia211/ice-pomelo/config"
	video "github.com/jiuxia211/ice-pomelo/kitex_gen/video"
	"github.com/jiuxia211/ice-pomelo/pkg/utils"
	"golang.org/x/sync/errgroup"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// UploadVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UploadVideo(ctx context.Context, req *video.UploadVideoRequest) (resp *video.UploadVideoResponse, err error) {
	resp = new(video.UploadVideoResponse)
	claims, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	var videoResp *db.Video
	nowTime := utils.GetTime()
	videoKey := fmt.Sprintf("%d_%s_video%s", claims.UserId, nowTime, req.VideoFormat)
	coverKey := fmt.Sprintf("%d_%s_cover%s", claims.UserId, nowTime, req.CoverFormat)
	videoURL := "https://" + config.ConfigInfo.Cos.BucketName + "." + config.ConfigInfo.Cos.Region + "/" + videoKey
	coverURL := "https://" + config.ConfigInfo.Cos.BucketName + "." + config.ConfigInfo.Cos.Region + "/" + coverKey

	var eg errgroup.Group

	// 上传视频
	eg.Go(func() error {
		err = service.NewVideoService(ctx).UploadVideo(req, videoKey)
		if err != nil {
			klog.Error(err)
			return err
		}
		return nil
	})
	// 上传封面
	eg.Go(func() error {
		err = service.NewVideoService(ctx).UploadCover(req, coverKey)
		if err != nil {
			klog.Error(err)
			return err
		}
		return nil
	})
	// 将视频数据写入数据库
	eg.Go(func() error {
		videoResp, err = service.NewVideoService(ctx).CreateVideo(req, videoURL, coverURL)
		if err != nil {
			klog.Error(err)
		}
		return err
	})
	if err := eg.Wait(); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Id = videoResp.Id
	return resp, nil
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)
	_, err = utils.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	videoList, err := service.NewVideoService(ctx).FeedVideo(req)
	if err != nil {
		return nil, err
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.VideoList(videoList)

	return resp, nil
}
