package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/video"
	"github.com/lukanzx/DouVo/kitex_gen/video/videoservice"
	"github.com/lukanzx/DouVo/pkg/constants"
	"github.com/lukanzx/DouVo/pkg/errno"
	"github.com/lukanzx/DouVo/pkg/middleware"
)

func InitVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetWorkCount(ctx context.Context, req *video.GetWorkCountRequest) (int64, error) {
	resp, err := videoClient.GetWorkCount(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.Base.Code != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.WorkCount, nil
}
