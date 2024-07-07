package main

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/lukanzx/DouVo/cmd/interaction/rpc"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"github.com/lukanzx/DouVo/kitex_gen/user"
)

func testCommentList(t *testing.T) {
	monkey.Patch(rpc.UserInfo, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: userId}, nil
	})

	defer monkey.UnpatchAll()

	req := &interaction.CommentListRequest{
		VideoId: videoId,
		Token:   token,
	}

	resp, err := interactionService.GetComments(req, 0)

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log(resp)
	t.Log("------------testCommentList success---------------")
}

func benchmarkCommentList(b *testing.B) {
	req := &interaction.CommentListRequest{
		VideoId: videoId,
		Token:   token,
	}

	for i := 0; i < b.N; i++ {
		_, err := interactionService.GetComments(req, 0)
		if err != nil {
			b.Error(err)
		}
	}
}
