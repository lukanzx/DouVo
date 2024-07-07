package main

import (
	"context"
	"testing"

	"github.com/lukanzx/DouVo/cmd/video/dal"
	"github.com/lukanzx/DouVo/cmd/video/dal/cache"
	"github.com/lukanzx/DouVo/cmd/video/rpc"
	"github.com/lukanzx/DouVo/cmd/video/service"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/pkg/utils"
)

var (
	videoService *service.VideoService
	videoId      []int64
	token        string
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()
	cache.Init()
	rpc.Init()
	videoService = service.NewVideoService(context.Background())
	token, _ = utils.CreateToken(10000)
	videoId = []int64{482581113097682944, 483299894140862464, 483302572409487360}
	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("Feed", testFeed)
	t.Run("Get publish", testGetPublishVideo)
	t.Run("Get favorite", testGetFavoriteVideo)
	t.Run("Get work count", testWorkCount)
	t.Run("Get video id by uid", testGetVideoIDByUid)
	t.Run("RPC Test", testRPC)
}
