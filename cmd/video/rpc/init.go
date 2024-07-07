package rpc

import (
	"github.com/lukanzx/DouVo/kitex_gen/interaction/interactionservice"
	"github.com/lukanzx/DouVo/kitex_gen/user/userservice"
)

var (
	userClient        userservice.Client
	interactionClient interactionservice.Client
)

func Init() {
	InitUserRPC()
	InitInteractionRPC()
}
