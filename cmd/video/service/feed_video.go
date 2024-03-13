package service

import (
	"github.com/jiuxia211/ice-pomelo/cmd/video/dal/db"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) (videoList []db.Video, err error) {
	videoList, err = db.GetVideoRandomly(s.ctx)
	if err != nil {
		return nil, err
	}

	return videoList, nil
}
