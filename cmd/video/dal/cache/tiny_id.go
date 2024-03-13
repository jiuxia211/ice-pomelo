package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
)

type TinyID struct {
	NextID int64
	MaxID  int64
}

func IsIDInfoExist(ctx context.Context) (exist int64, err error) {
	exist, err = RedisClient.Exists(ctx, strconv.FormatInt(constants.VideoBizType, 10)).Result()
	klog.Debugf("redis biz type %v id info exist status: %v\n", constants.VideoBizType, exist)
	return
}

func SetID(ctx context.Context, tinyID *TinyID) {
	idJson, err := json.Marshal(&tinyID)
	if err != nil {
		klog.Error(err)
	}
	err = RedisClient.Set(ctx, strconv.FormatInt(constants.VideoBizType, 10), idJson, 0).Err()
	if err != nil {
		klog.Error(err)
	}
	klog.Debugf("redis biz type %v set id info %+v\n", constants.VideoBizType, tinyID)
}

func GetID(ctx context.Context) (tinyID *TinyID, err error) {
	data, err := RedisClient.Get(ctx, strconv.FormatInt(constants.VideoBizType, 10)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &tinyID)

	if err != nil {
		return nil, err
	}
	klog.Debugf("redis biz type %v get id info %+v\n", constants.VideoBizType, tinyID)
	return tinyID, err
}
