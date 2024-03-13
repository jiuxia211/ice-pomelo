package service

import (
	"context"

	"github.com/jiuxia211/ice-pomelo/cmd/video/dal/cache"
	"github.com/jiuxia211/ice-pomelo/cmd/video/rpc"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func getID(ctx context.Context) (id int64, err error) {
	tinyID := new(cache.TinyID)

	tinyID, err = cache.GetID(ctx)
	if err != nil {
		return 0, err
	}

	if tinyID.NextID > tinyID.MaxID {
		maxID, err := rpc.GetMaxTinyID(ctx, &tiny_id.GetMaxIDRequest{BizType: constants.VideoBizType})
		if err != nil {
			return 0, err
		}
		// maxID+1直接给当前数据
		id = maxID + 1
		go cache.SetID(ctx, &cache.TinyID{
			NextID: maxID + 2,
			MaxID:  maxID + constants.IDStep,
		})

	} else {
		id = tinyID.NextID
		go cache.SetID(ctx, &cache.TinyID{
			NextID: tinyID.NextID + 1,
			MaxID:  tinyID.MaxID,
		})
	}

	return id, nil
}
