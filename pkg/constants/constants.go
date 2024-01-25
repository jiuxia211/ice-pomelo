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

	UserRedisDB = 1

	UserBizType = 1000
	JWTValue    = "ice_pomelo"

	UserServiceName   = "user"
	TinyIDServiceName = "tiny_id"
)
