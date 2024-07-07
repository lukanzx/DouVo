package test

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/lukanzx/DouVo/kitex_gen/chat/messageservice"
	"github.com/lukanzx/DouVo/pkg/constants"
)

func testRpc(t *testing.T) {
	_, err := messageservice.NewClient("chat",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:10003"))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
