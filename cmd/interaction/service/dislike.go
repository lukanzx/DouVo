package service

import (
	"errors"

	"github.com/lukanzx/DouVo/cmd/interaction/dal/cache"
	"github.com/lukanzx/DouVo/cmd/interaction/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"github.com/lukanzx/DouVo/pkg/errno"
	"gorm.io/gorm"
)

func (s *InteractionService) Dislike(req *interaction.FavoriteActionRequest, userID int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userID)
	if err != nil {
		return err
	}
	if !exist {
		err := db.IsFavorited(s.ctx, userID, req.VideoId, 1)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LikeNoExistError
		}
	}

	ok, _, err := cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		return err
	}
	if ok {
		if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId, userID); err != nil {
			return err
		}
	}

	// write into mysql
	return db.UpdateFavoriteStatus(s.ctx, userID, req.VideoId, 0)
}
