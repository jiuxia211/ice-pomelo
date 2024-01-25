package cache

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

func SetCode(ctx context.Context, email string, code string) {
	err := RedisClient.Set(ctx, email, code, 10*time.Minute).Err()
	if err != nil {
		klog.Error(err)
	}
}
func CheckCode(ctx context.Context, email string, code string) (success bool, err error) {
	data, err := RedisClient.Get(ctx, email).Result()
	if err != nil {
		return false, err
	}
	return data == code, err

}
