package main

import (
	"testing"

	"github.com/lukanzx/DouVo/kitex_gen/user"
)

func testRegister(t *testing.T) {
	_, err := userService.CreateUser(&user.RegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
