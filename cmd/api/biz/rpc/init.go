package rpc

import (
	"github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id/tinyidservice"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user/userservice"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video/videoservice"
)

var (
	tinyIDClient tinyidservice.Client
	userClient   userservice.Client
	videoClient  videoservice.Client
)

func Init() {
	InitTinyIDRPC()
	InitUserRPC()
	InitVideoRPC()
}
