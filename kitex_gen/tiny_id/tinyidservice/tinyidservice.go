// Code generated by Kitex v0.8.0. DO NOT EDIT.

package tinyidservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	tiny_id "github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id"
)

func serviceInfo() *kitex.ServiceInfo {
	return tinyIDServiceServiceInfo
}

var tinyIDServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TinyIDService"
	handlerType := (*tiny_id.TinyIDService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetMaxID": kitex.NewMethodInfo(getMaxIDHandler, newTinyIDServiceGetMaxIDArgs, newTinyIDServiceGetMaxIDResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "tiny_id",
		"ServiceFilePath": `../../idl/tiny_id.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func getMaxIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*tiny_id.TinyIDServiceGetMaxIDArgs)
	realResult := result.(*tiny_id.TinyIDServiceGetMaxIDResult)
	success, err := handler.(tiny_id.TinyIDService).GetMaxID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTinyIDServiceGetMaxIDArgs() interface{} {
	return tiny_id.NewTinyIDServiceGetMaxIDArgs()
}

func newTinyIDServiceGetMaxIDResult() interface{} {
	return tiny_id.NewTinyIDServiceGetMaxIDResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetMaxID(ctx context.Context, req *tiny_id.GetMaxIDRequest) (r *tiny_id.GetMaxIDResponse, err error) {
	var _args tiny_id.TinyIDServiceGetMaxIDArgs
	_args.Req = req
	var _result tiny_id.TinyIDServiceGetMaxIDResult
	if err = p.c.Call(ctx, "GetMaxID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
