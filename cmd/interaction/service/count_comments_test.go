package service

import (
	"context"
	"testing"

	"github.com/lukanzx/DouVo/cmd/interaction/dal"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"github.com/lukanzx/DouVo/pkg/utils"
)

var (
	userId      int64
	videoId     int64
	token       string
	commentText string
	commentId   int64

	interactionService *InteractionService
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init("../../../config")
	interactionService = NewInteractionService(context.Background())

	userId = 10000
	token, _ = utils.CreateToken(userId)
	commentText = "发条评论看看"

	videoId = 1
	m.Run()
}

func TestCommentCount(t *testing.T) {
	req := &interaction.CommentCountRequest{
		VideoId: 1,
		Token:   &token,
	}
	_, err := interactionService.CountComments(req, 0)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
