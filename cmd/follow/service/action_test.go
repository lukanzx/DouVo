package service_test

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lukanzx/DouVo/cmd/follow/dal"
	"github.com/lukanzx/DouVo/cmd/follow/rpc"
	"github.com/lukanzx/DouVo/cmd/follow/service"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/follow"
	"github.com/lukanzx/DouVo/pkg/utils"
)

type Test struct {
	id         int64
	touserid   int64
	token      string
	actiontype int64
}

var actionTests = []Test{
	{10001, 10002, "", 2},
	{10002, 10001, "", 2},
	{10001, 10002, "", 1},
	{10001, 10002, "", 1},
	{10001, 10002, "", 2},
	{10001, 10002, "", 2},
	{10001, 10002, "", 1},
	{10001, 10002, "hhh", 1},
	{10001, 10002, "", 3},
	{10002, 10086, "", 1},
	{10001, 10001, "", 1},
	{10001, 10002, "", 1},
	{10002, 10001, "", 1},
}

func TestAction(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	followService := service.NewFollowService(context.Background())
	for i, test := range actionTests {
		test.token, _ = utils.CreateToken(test.id)
		err := followService.Action(&follow.ActionRequest{
			Token:      test.token,
			ToUserId:   test.touserid,
			ActionType: test.actiontype,
		})

		if err != nil {
			klog.Infof("test num %v,err:%v", i, err)
			continue
		}
	}
}
