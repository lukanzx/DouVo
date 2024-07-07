package dal

import (
	"github.com/lukanzx/DouVo/cmd/chat/dal/cache"
	"github.com/lukanzx/DouVo/cmd/chat/dal/db"
	"github.com/lukanzx/DouVo/cmd/chat/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.InitRabbitMQ()
	mq.InitChatMQ()
}
