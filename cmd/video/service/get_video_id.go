package service

import (
	"github.com/lukanzx/DouVo/cmd/video/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/video"
)

func (s *VideoService) GetVideoIDByUid(req *video.GetVideoIDByUidRequset) (videoIDList []int64, err error) {
	return db.GetVideoIDByUid(s.ctx, req.UserId)
}
