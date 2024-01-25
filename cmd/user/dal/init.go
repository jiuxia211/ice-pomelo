package dal

import (
	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/cache"
	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
