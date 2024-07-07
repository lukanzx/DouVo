package rpc

import (
	"github.com/lukanzx/DouVo/kitex_gen/user/userservice"
	"github.com/lukanzx/DouVo/kitex_gen/video/videoservice"
)

var (
	userClient  userservice.Client
	videoClient videoservice.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
}
