package service

import (
	"github.com/lukanzx/DouVo/cmd/follow/dal/cache"
	"github.com/lukanzx/DouVo/cmd/follow/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/follow"
)

func (s *FollowService) FollowCount(req *follow.FollowCountRequest) (int64, error) {
	// 先进入redis中查询
	followCount, err := cache.FollowCount(s.ctx, req.UserId)
	if err != nil {
		return -1, err
	}

	if followCount == 0 { // redis中没查到,进入db中查
		followCount, err = db.FollowCount(s.ctx, req.UserId)
	}
	if err != nil {
		return -1, err
	}
	return followCount, nil
}
