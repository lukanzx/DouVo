package service_test

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lukanzx/DouVo/cmd/follow/dal"
	"github.com/lukanzx/DouVo/cmd/follow/service"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/follow"
	"github.com/lukanzx/DouVo/pkg/utils"
)

var followerCountTests = []Test{
	{10001, 10002, "", 1},
	{001, 002, "", 1},
}

func TestFollowerCount(t *testing.T) {
	config.InitForTest()
	dal.Init()
	followService := service.NewFollowService(context.Background())
	for i, test := range followerCountTests {
		test.token, _ = utils.CreateToken(test.id)
		_, err := followService.FollowerCount(&follow.FollowerCountRequest{
			UserId: test.id,
			Token:  test.token,
		})

		if err != nil {
			klog.Infof("test num %v,err:%v", i, err)
			continue
		}
	}
}
