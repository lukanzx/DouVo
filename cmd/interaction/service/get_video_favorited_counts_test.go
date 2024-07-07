package service

import (
	"testing"

	"github.com/lukanzx/DouVo/kitex_gen/interaction"
)

func TestGetVideoFavoritedCount(t *testing.T) {
	req := &interaction.VideoFavoritedCountRequest{
		VideoId: videoId,
		Token:   token,
	}

	_, err := interactionService.GetVideoFavoritedCount(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
