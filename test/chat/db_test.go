package test

import (
	"testing"

	"github.com/lukanzx/DouVo/cmd/chat/dal/db"
)

func testDB(t *testing.T) {
	db.Init()
	if db.DB == nil {
		t.Fail()
	}
}
