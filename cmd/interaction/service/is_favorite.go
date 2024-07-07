package service

import (
	"errors"

	"github.com/lukanzx/DouVo/cmd/interaction/dal/cache"
	"github.com/lukanzx/DouVo/cmd/interaction/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"gorm.io/gorm"
)

func (s *InteractionService) IsFavorite(req *interaction.IsFavoriteRequest) (bool, error) {
	// read from redis
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, req.UserId)
	if err != nil {
		return exist, err
	}
	if exist {
		return exist, nil
	}
	// read from mysql
	err = db.IsFavorited(s.ctx, req.UserId, req.VideoId, 1)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
