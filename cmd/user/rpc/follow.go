package rpc

import (
	"context"
	"fmt"
	"github.com/lukanzx/DouVo/config"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/lukanzx/DouVo/kitex_gen/follow"
	"github.com/lukanzx/DouVo/kitex_gen/follow/followservice"
	"github.com/lukanzx/DouVo/pkg/constants"
	"github.com/lukanzx/DouVo/pkg/errno"
	"github.com/lukanzx/DouVo/pkg/middleware"
)

func InitFollowRPC() {

	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	//r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})

	if err != nil {
		fmt.Println(" ... ")
		panic(err)
	}
	c, err := followservice.NewClient(
		constants.FollowServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)
	if err != nil { // 如果创建过程中出现错误，立即终止程序并输出错误信息
		panic(err)
	}
	followClient = c
}

func GetFollowCount(ctx context.Context, req *follow.FollowCountRequest) (int64, error) {
	resp, err := followClient.FollowCount(ctx, req)
	if err != nil { // 如果调用过程中出现错误，则返回错误信息
		return -1, err
	}
	if resp.Base.Code != errno.SuccessCode { // 检查响应中的错误码，如果不是成功的响应，则返回对应的错误信息
		return -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return *resp.FollowCount, nil
}

func GetFollowerCount(ctx context.Context, req *follow.FollowerCountRequest) (int64, error) {
	resp, err := followClient.FollowerCount(ctx, req)
	if err != nil { // 如果调用过程中出现错误，则返回错误信息
		return -1, err
	}
	if resp.Base.Code != errno.SuccessCode { // 检查响应中的错误码，如果不是成功的响应，则返回对应的错误信息
		return -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return *resp.FollowerCount, nil
}

func IsFollow(ctx context.Context, req *follow.IsFollowRequest) (bool, error) {
	resp, err := followClient.IsFollow(ctx, req)
	if err != nil { // 如果调用过程中出现错误，则返回错误信息
		return false, err
	}
	if resp.Base.Code != errno.SuccessCode { // 检查响应中的错误码，如果不是成功的响应，则返回对应的错误信息
		return false, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.IsFollow, nil
}
