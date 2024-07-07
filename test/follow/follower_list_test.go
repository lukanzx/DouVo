package main

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/lukanzx/DouVo/cmd/follow/rpc"
	"github.com/lukanzx/DouVo/kitex_gen/follow"
	"github.com/lukanzx/DouVo/kitex_gen/user"
)

func testFollowerList(t *testing.T) {
	monkey.Patch(rpc.GetUser, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: touserid}, nil
	})

	defer monkey.UnpatchAll()
	_, err := followService.FollowerList(&follow.FollowerListRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func benchmarkFollowerList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := followService.FollowerList(&follow.FollowerListRequest{
			UserId: id,
			Token:  token,
		})

		if err != nil {
			b.Errorf("err: [%v] \n", err)
		}

		time.Sleep(100 * time.Millisecond) // Add a sleep to simulate some processing time
	}
}
