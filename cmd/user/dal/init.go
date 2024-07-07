package dal

import (
	"github.com/lukanzx/DouVo/cmd/user/dal/cache"
	"github.com/lukanzx/DouVo/cmd/user/dal/db"
	"github.com/lukanzx/DouVo/cmd/user/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.Init()
}
