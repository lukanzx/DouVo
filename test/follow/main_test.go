package main

import (
	"context"
	"testing"

	"github.com/lukanzx/DouVo/cmd/follow/dal"
	"github.com/lukanzx/DouVo/cmd/follow/rpc"
	"github.com/lukanzx/DouVo/cmd/follow/service"
	"github.com/lukanzx/DouVo/config"
)

var (
	touserid   int64
	actiontype int64
	token      string
	id         int64

	followService *service.FollowService
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()
	rpc.Init()

	followService = service.NewFollowService(context.Background())

	touserid = 10002
	actiontype = 1
	id = 10001

	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("action", testAction)

	t.Run("followList", testFollowList)

	t.Run("followerList", testFollowerList)

	t.Run("friendList", testFriendList)

	t.Run("followCount", testFollowCount)

	t.Run("followerCount", testFollowerCount)

	t.Run("isfollow", testIsFollow)

	t.Run("RPC Test", testRPC)
}

func BenchmarkMainOrder(b *testing.B) {
	b.Run("action", benchmarkAction)

	b.Run("followList", benchmarkFollowList)

	b.Run("followerList", benchmarkFollowerList)

	b.Run("friendList", benchmarkFriendList)
}
