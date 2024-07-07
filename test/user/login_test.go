package main

import (
	"testing"

	"github.com/lukanzx/DouVo/kitex_gen/user"
	"github.com/lukanzx/DouVo/pkg/utils"
)

func testLogin(t *testing.T) {
	resp, err := userService.CheckUser(&user.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	token, err = utils.CreateToken(resp.Id)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	id = resp.Id
}
