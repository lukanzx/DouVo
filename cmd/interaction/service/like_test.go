package service

import (
	"testing"

	"github.com/lukanzx/DouVo/kitex_gen/interaction"
)

func TestLike(t *testing.T) {
	req := &interaction.FavoriteActionRequest{
		VideoId:    videoId,
		Token:      token,
		ActionType: 1,
	}

	err := interactionService.Like(req, userId)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
