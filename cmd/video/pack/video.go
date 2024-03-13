package pack

import (
	"github.com/jiuxia211/ice-pomelo/cmd/video/dal/db"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video"
)

func Video(data *db.Video) *video.Video {
	if data == nil {
		return nil
	}
	return &video.Video{
		Id:           data.Id,
		Uid:          data.UserID,
		VideoUrl:     data.VideoUrl,
		CoverUrl:     data.CoverUrl,
		Title:        data.Title,
		Introduction: data.Introduction,
	}
}
func VideoList(data []db.Video) []*video.Video {
	videoList := make([]*video.Video, 0, len(data))
	for i := 0; i < len(data); i++ {
		videoList = append(videoList, Video(&data[i]))
	}
	return videoList
}
