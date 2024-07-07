package main

import (
	"fmt"
	"testing"

	"github.com/lukanzx/DouVo/cmd/video/pack"
	"github.com/lukanzx/DouVo/kitex_gen/video"
)

func testFeed(t *testing.T) {
	testTime := new(int64)
	*testTime = 1693101739
	videoList, userList, favoriteCountList, commentCountList, isFavoriteList, err := videoService.FeedVideo(&video.FeedRequest{
		LatestTime: testTime,
		Token:      &token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(pack.VideoList(videoList, userList, favoriteCountList, commentCountList, isFavoriteList))
}
