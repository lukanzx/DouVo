package service_test

import (
	"context"
	"testing"

	"github.com/lukanzx/DouVo/cmd/video/dal"
	"github.com/lukanzx/DouVo/cmd/video/rpc"
	"github.com/lukanzx/DouVo/cmd/video/service"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/video"
	"github.com/lukanzx/DouVo/pkg/utils"
)

func TestGetWorkCount(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	token, err := utils.CreateToken(10000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	videoService := service.NewVideoService(context.Background())
	_, err = videoService.GetWorkCount(&video.GetWorkCountRequest{Token: token, UserId: 10000})

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
