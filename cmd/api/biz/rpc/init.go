package rpc

import (
	"github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id/tinyidservice"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user/userservice"
)

var (
	tinyIDClient tinyidservice.Client
	userClient   userservice.Client
)

func Init() {
	InitTinyIDRPC()
	InitUserRPC()
}
