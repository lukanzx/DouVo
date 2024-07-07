package service

import (
	"github.com/lukanzx/DouVo/cmd/user/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser create user info
func (s *UserService) CreateUser(req *user.RegisterRequest) (*db.User, error) {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	userModel := &db.User{
		Username: req.Username,
		Password: string(hashBytes),
	}

	return db.CreateUser(s.ctx, userModel)
}
