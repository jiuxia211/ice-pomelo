// Code generated by hertz generator.

package api

import (
	"context"
	"io"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/jiuxia211/ice-pomelo/cmd/api/biz/model/api"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/pack"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/rpc"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video"
)

// Feed .
// @router /pomelo/video/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.FeedResponse)

	videoList, err := rpc.VideoFeed(ctx, &video.FeedRequest{
		Token: req.Token,
	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.VideoList(videoList)
	pack.SendResponse(c, resp)
}

// UploadVideo .
// @router /pomelo/video/upload [POST]
func UploadVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UploadVideoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	videoFile, err := c.FormFile("videoFile")
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	// 简单的使用文件后缀来获取图片格式
	videoFileExt := filepath.Ext(videoFile.Filename)
	videoFileContent, err := videoFile.Open()
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	videoByteContainer, err := io.ReadAll(videoFileContent)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	coverFile, err := c.FormFile("coverFile")
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	coverFileExt := filepath.Ext(coverFile.Filename)
	coverFileContent, err := coverFile.Open()
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	coverByteContainer, err := io.ReadAll(coverFileContent)
	resp := new(api.UploadVideoResponse)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	videoID, err := rpc.VideoUpload(ctx, &video.UploadVideoRequest{
		Token:        req.Token,
		VideoFile:    videoByteContainer,
		VideoFormat:  videoFileExt,
		CoverFile:    coverByteContainer,
		CoverFormat:  coverFileExt,
		Title:        req.Title,
		Introduction: req.Introduction,
	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.ID = videoID
	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}
