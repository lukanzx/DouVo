package main

import (
	"testing"

	"github.com/lukanzx/DouVo/kitex_gen/follow"
)

func testIsFollow(t *testing.T) {
	_, err := followService.IsFollow(&follow.IsFollowRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
