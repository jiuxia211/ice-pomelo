package rpc

import "github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id/tinyidservice"

var (
	tinyIDClient tinyidservice.Client
)

func Init() {
	InitTinyIDRPC()
}
