package rpc

import (
	"github.com/lukanzx/DouVo/kitex_gen/chat/messageservice"
	"github.com/lukanzx/DouVo/kitex_gen/user/userservice"
)

var (
	userClient userservice.Client
	chatClient messageservice.Client
)

func Init() {
	InitUserRPC()
	InitChatRPC()
}
