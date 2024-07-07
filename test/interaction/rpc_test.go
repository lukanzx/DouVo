package main

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/lukanzx/DouVo/kitex_gen/interaction/interactionservice"
	"github.com/lukanzx/DouVo/pkg/constants"
)

func testRPC(t *testing.T) {
	_, err := interactionservice.NewClient("interaction",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:10005"))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
