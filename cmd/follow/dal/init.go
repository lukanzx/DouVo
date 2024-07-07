package dal

import (
	"github.com/lukanzx/DouVo/cmd/follow/dal/cache"
	"github.com/lukanzx/DouVo/cmd/follow/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
