package service

import (
	"testing"
	"time"

	"github.com/lukanzx/DouVo/kitex_gen/chat"
	"github.com/lukanzx/DouVo/pkg/utils"
)

func testPostMessage(t *testing.T) {
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req := &chat.MessagePostRequest{
		Token:      token,
		ToUserId:   to_user_id,
		Content:    content,
		ActionType: &ac_type,
	}
	err = chatservice.SendMessage(req, from_user_id, create_at)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req.ToUserId = from_user_id
	err = chatservice.SendMessage(req, to_user_id, create_at)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	time.Sleep(2 * time.Second)
}
