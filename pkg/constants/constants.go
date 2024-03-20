package constants

import "time"

const (
	MaxIdleConns    = 10
	MaxConnections  = 100
	ConnMaxLifetime = time.Hour

	IDStep  = 1000
	StartID = 1000

	UserTableName   = "user"
	TinyIDTableName = "tiny_id"
	VideoTableName  = "video"
	ChatTableName   = "message"

	UserRedisDB  = 1
	VideoRedisDB = 2

	UserBizType  = 1000
	VideoBizType = 2000
	JWTValue     = "ice_pomelo"

	UserServiceName   = "user"
	TinyIDServiceName = "tiny_id"
	VideoServiceName  = "video"
	ChatServiceName   = "chat"
)
