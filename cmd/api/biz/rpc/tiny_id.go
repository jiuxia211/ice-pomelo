package rpc

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id/tinyidservice"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

func InitTinyIDRPC() {
	c, err := tinyidservice.NewClient(
		constants.TinyIDServiceName,
		client.WithHostPorts("127.0.0.1:8210"))
	if err != nil {
		klog.Fatalf("tiny id rpc init error: %v", err)
	}
	tinyIDClient = c
}

func GetMaxTinyID(ctx context.Context, req *tiny_id.GetMaxIDRequest) (maxID int64, err error) {
	resp, err := tinyIDClient.GetMaxID(ctx, req)

	if err != nil {
		return 0, err
	}

	if resp.Base.Code != errz.SuccessCode {
		return 0, errors.New(resp.Base.Msg)
	}

	return resp.MaxID, nil
}
