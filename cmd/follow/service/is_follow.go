package service

import (
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lukanzx/DouVo/cmd/follow/dal/cache"
	"github.com/lukanzx/DouVo/cmd/follow/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/follow"
)

func (s *FollowService) IsFollow(req *follow.IsFollowRequest) (bool, error) {
	// 先进入redis中判断是否有关注
	ex1, err := cache.IsFollow(s.ctx, req.UserId, req.ToUserId)

	if err != nil {
		return false, err
	}

	if ex1 {
		return true, nil
	}

	ex2, err := db.IsFollow(s.ctx, req.UserId, req.ToUserId)

	if err != nil {
		if errors.Is(err, db.RecordNotFound) {
			return false, nil
		}

		klog.Errorf("db sql meet error: %v\n", err)
		return false, err
	}

	return ex2, nil
}
