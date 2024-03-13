package pack

import (
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/model/api"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video"
)

func Video(data *video.Video) *api.Video {
	if data == nil {
		return nil
	}
	return &api.Video{
		ID:           data.Id,
		UID:          data.Uid,
		VideoURL:     data.VideoUrl,
		CoverURL:     data.CoverUrl,
		Title:        data.Title,
		Introduction: data.Introduction,
	}
}
func VideoList(data []*video.Video) []*api.Video {
	resp := make([]*api.Video, 0)
	for i := 0; i < len(data); i++ {
		resp = append(resp, Video(data[i]))
	}
	return resp
}
