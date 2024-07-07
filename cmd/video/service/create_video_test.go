package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lukanzx/DouVo/cmd/video/dal"
	"github.com/lukanzx/DouVo/cmd/video/rpc"
	"github.com/lukanzx/DouVo/cmd/video/service"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/video"
	"github.com/lukanzx/DouVo/pkg/errno"
	"github.com/lukanzx/DouVo/pkg/utils"
)

func TestCreateVideo(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	videoService := service.NewVideoService(context.Background())

	_, err := videoService.CreateVideo(&video.PutVideoRequest{VideoFile: nil,
		Title: "test_title",
		Token: "",
	}, "test_video_URL", "test_cover_URL")
	if !errors.Is(err, errno.AuthorizationFailedError) {
		t.Error(err)
		t.Fail()
	}

	token, err := utils.CreateToken(10000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	_, err = videoService.CreateVideo(&video.PutVideoRequest{VideoFile: nil,
		Title: "test_title",
		Token: token,
	}, "test_video_URL", "test_cover_URL")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
