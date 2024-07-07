package rpc

import (
	"github.com/lukanzx/DouVo/kitex_gen/follow/followservice"
	"github.com/lukanzx/DouVo/kitex_gen/interaction/interactionservice"
	"github.com/lukanzx/DouVo/kitex_gen/video/videoservice"
)

var (
	followClient      followservice.Client
	interactionClient interactionservice.Client
	videoClient       videoservice.Client
)

func Init() {
	InitFollowRPC()
	InitInteractionRPC()
	InitVideoRPC()
}
