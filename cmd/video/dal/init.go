package dal

import (
	"github.com/jiuxia211/ice-pomelo/cmd/video/dal/cache"
	"github.com/jiuxia211/ice-pomelo/cmd/video/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
