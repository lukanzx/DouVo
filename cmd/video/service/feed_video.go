package service

import (
	"time"

	"github.com/lukanzx/DouVo/cmd/video/dal/cache"
	"github.com/lukanzx/DouVo/cmd/video/dal/db"
	"github.com/lukanzx/DouVo/cmd/video/rpc"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"github.com/lukanzx/DouVo/kitex_gen/user"
	"github.com/lukanzx/DouVo/kitex_gen/video"
	"github.com/lukanzx/DouVo/pkg/utils"
	"golang.org/x/sync/errgroup"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, []*user.User, []int64, []int64, []bool, error) {
	var videoList []db.Video
	var err error
	var claims *utils.Claims

	if claims, err = utils.CheckToken(*req.Token); err != nil {
		return nil, nil, nil, nil, nil, err
	}

	if exist, err := cache.IsExistVideoInfo(s.ctx, *req.LatestTime); exist == 1 {
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		videoList, err = cache.GetVideoList(s.ctx, *req.LatestTime)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
	} else {
		formattedTime := time.UnixMilli(*req.LatestTime).Format("2006-01-02 15:04:05")
		videoList, err = db.GetVideoInfoByTime(s.ctx, formattedTime)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		go cache.AddVideoList(s.ctx, videoList, *req.LatestTime)
	}
	var eg errgroup.Group
	type result struct {
		userInfo      *user.User
		favoriteCount int64
		commentCount  int64
		isFavorite    bool
	}
	results := make([]result, len(videoList))

	for i := 0; i < len(videoList); i++ {
		index := i
		eg.Go(func() error {

			userInfo, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: videoList[index].UserID,
				Token:  *req.Token,
			})
			if err != nil {
				return err
			}

			favoriteCount, err := rpc.GetVideoFavoriteCount(s.ctx, &interaction.VideoFavoritedCountRequest{
				VideoId: videoList[index].Id,
				Token:   *req.Token,
			})
			if err != nil {
				return err
			}

			commentCount, err := rpc.GetCommentCount(s.ctx, &interaction.CommentCountRequest{
				VideoId: videoList[index].Id,
				Token:   req.Token,
			})
			if err != nil {
				return err
			}
			isFavorite, err := rpc.GetVideoIsFavorite(s.ctx, &interaction.InteractionServiceIsFavoriteArgs{Req: &interaction.IsFavoriteRequest{
				UserId:  claims.UserId,
				VideoId: videoList[index].Id,
				Token:   *req.Token,
			}})
			if err != nil {
				return err
			}
			results[index] = result{
				userInfo:      userInfo,
				favoriteCount: favoriteCount,
				commentCount:  commentCount,
				isFavorite:    isFavorite,
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, nil, nil, nil, nil, err
	}
	var userList []*user.User
	var favoriteCountList []int64
	var commentCountList []int64
	var isFavoriteList []bool

	for _, result := range results {
		userList = append(userList, result.userInfo)
		favoriteCountList = append(favoriteCountList, result.favoriteCount)
		commentCountList = append(commentCountList, result.commentCount)
		isFavoriteList = append(isFavoriteList, result.isFavorite)
	}

	return videoList, userList, favoriteCountList, commentCountList, isFavoriteList, err
}
