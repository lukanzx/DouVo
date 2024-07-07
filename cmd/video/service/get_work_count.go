package service

import (
	"github.com/lukanzx/DouVo/cmd/video/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/video"
)

func (s *VideoService) GetWorkCount(req *video.GetWorkCountRequest) (workCount int64, err error) {
	return db.GetWorkCountByUid(s.ctx, req.UserId)
}
