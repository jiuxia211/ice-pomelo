package main

import (
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/jiuxia211/ice-pomelo/cmd/user/dal"
	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/cache"
	"github.com/jiuxia211/ice-pomelo/cmd/user/rpc"
	"github.com/jiuxia211/ice-pomelo/config"
	user "github.com/jiuxia211/ice-pomelo/kitex_gen/user/userservice"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
)

func tinyIDInit() {
	exist, err := cache.IsIDInfoExist(context.Background())
	if err != nil {
		klog.Fatalf("tinyID init error: %v", err)
	}
	if exist == 0 {
		cache.SetID(context.Background(), &cache.TinyID{
			NextID: constants.StartID,
			MaxID:  0,
		})
	}
}

func main() {
	config.Init()
	dal.Init()
	rpc.Init()
	tinyIDInit()
	klog.SetLevel(klog.LevelDebug)
	// TODO
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8211")
	if err != nil {
		klog.Fatalf("resolve addr error: %v", err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServiceName,
		}),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
