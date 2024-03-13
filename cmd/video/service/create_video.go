package service

import (
	"github.com/jiuxia211/ice-pomelo/cmd/video/dal/db"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video"
	"github.com/jiuxia211/ice-pomelo/pkg/utils"
)

func (s *VideoService) CreateVideo(req *video.UploadVideoRequest, videoURL string, coverURL string) (*db.Video, error) {
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}
	id, err := getID(s.ctx)
	if err != nil {
		return nil, err
	}

	videoModel := &db.Video{
		Id:           id,
		UserID:       claim.UserId,
		VideoUrl:     videoURL,
		CoverUrl:     coverURL,
		Title:        req.Title,
		Introduction: req.Introduction,
	}
	return db.CreateVideo(s.ctx, videoModel)
}
