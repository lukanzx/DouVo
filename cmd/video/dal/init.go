package dal

import (
	"github.com/lukanzx/DouVo/cmd/video/dal/cache"
	"github.com/lukanzx/DouVo/cmd/video/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
