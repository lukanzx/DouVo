package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/video"
	"github.com/lukanzx/DouVo/kitex_gen/video/videoservice"
	"github.com/lukanzx/DouVo/pkg/constants"
	"github.com/lukanzx/DouVo/pkg/errno"
	"github.com/lukanzx/DouVo/pkg/middleware"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

func InitVideoRPC() {
	resolver, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	client, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(resolver),
		client.WithSuite(trace.NewDefaultClientSuite()),
	)

	if err != nil {
		panic(err)
	}

	videoClient = client
}

func GetFavoriteVideoList(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetFavoriteVideoInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.VideoList, nil
}

func GetUserVideoList(ctx context.Context, req *video.GetVideoIDByUidRequset) ([]int64, error) {
	resp, err := videoClient.GetVideoIDByUid(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.VideoId, nil
}
