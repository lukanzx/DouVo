package main

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/lukanzx/DouVo/kitex_gen/follow/followservice"
	"github.com/lukanzx/DouVo/pkg/constants"
)

func testRPC(t *testing.T) {
	_, err := followservice.NewClient("follow",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:10004"))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
