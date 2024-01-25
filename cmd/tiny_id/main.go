package main

import (
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/jiuxia211/ice-pomelo/cmd/tiny_id/dal"
	"github.com/jiuxia211/ice-pomelo/cmd/tiny_id/service"
	"github.com/jiuxia211/ice-pomelo/config"
	tiny_id "github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id/tinyidservice"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
)

func bizInit() {
	service.NewTinyIDService(context.Background()).CreateTinyID(constants.UserBizType)
}

func main() {
	config.Init()
	dal.Init()
	bizInit()
	klog.SetLevel(klog.LevelDebug)
	// TODO
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8210")
	if err != nil {
		klog.Fatalf("resolve addr error: %v", err)
	}

	svr := tiny_id.NewServer(
		new(TinyIDServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.TinyIDServiceName,
		}),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
