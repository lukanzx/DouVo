package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/lukanzx/DouVo/config"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"github.com/lukanzx/DouVo/kitex_gen/interaction/interactionservice"
	"github.com/lukanzx/DouVo/pkg/constants"
	"github.com/lukanzx/DouVo/pkg/errno"
	"github.com/lukanzx/DouVo/pkg/middleware"
)

func InitInteractionRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	c, err := interactionservice.NewClient(
		constants.InteractionServiceName,
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
	interactionClient = c
}

func GetFavoriteCount(ctx context.Context, req *interaction.UserFavoriteCountRequest) (int64, error) {
	resp, err := interactionClient.UserFavoriteCount(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.Base.Code != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}
	return resp.LikeCount, nil
}

func GetTotalFavorited(ctx context.Context, req *interaction.UserTotalFavoritedRequest) (int64, error) {
	resp, err := interactionClient.UserTotalFavorited(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.Base.Code != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}
	return resp.TotalFavorited, nil
}
