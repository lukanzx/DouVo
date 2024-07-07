package db

import (
	"context"
	"errors"
	"time"

	"github.com/lukanzx/DouVo/pkg/errno"
	"gorm.io/gorm"
)

type User struct {
	Id              int64
	Username        string
	Password        string
	Avatar          string `gorm:"default:https://th.bing.com/th/id/OIP.Tap7BMzkb9-0-gz9Q0bijQAAAA?w=139&h=180&c=7&r=0&o=5&dpr=1.3&pid=1.7"`
	BackgroundImage string `gorm:"default:https://th.bing.com/th/id/OIP.sw69HsYv8yzq1I1bc3HGLwAAAA?w=229&h=180&c=7&r=0&o=5&dpr=1.3&pid=1.7"`
	Signature       string `gorm:"default:NOT NULL BUT SEEMS NULL"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func CreateUser(ctx context.Context, user *User) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", user.Username).First(&userResp).Error

	if err == nil {
		return nil, errno.UserExistedError
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := DB.WithContext(ctx).Create(user).Error; err != nil {
		// add some logs
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", username).First(&userResp).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return userResp, nil
}

func GetUserByID(ctx context.Context, userid int64) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("id = ?", userid).First(&userResp).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFoundError
		}
		return nil, err
	}

	return userResp, nil
}
