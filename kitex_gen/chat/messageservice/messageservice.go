// Code generated by Kitex v0.6.2. DO NOT EDIT.

package messageservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	chat "github.com/lukanzx/DouVo/kitex_gen/chat"
)

func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

var messageServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MessageService"
	handlerType := (*chat.MessageService)(nil)
	methods := map[string]kitex.MethodInfo{
		"MessagePost": kitex.NewMethodInfo(messagePostHandler, newMessageServiceMessagePostArgs, newMessageServiceMessagePostResult, false),
		"MessageList": kitex.NewMethodInfo(messageListHandler, newMessageServiceMessageListArgs, newMessageServiceMessageListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "chat",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func messagePostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*chat.MessageServiceMessagePostArgs)
	realResult := result.(*chat.MessageServiceMessagePostResult)
	success, err := handler.(chat.MessageService).MessagePost(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceMessagePostArgs() interface{} {
	return chat.NewMessageServiceMessagePostArgs()
}

func newMessageServiceMessagePostResult() interface{} {
	return chat.NewMessageServiceMessagePostResult()
}

func messageListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*chat.MessageServiceMessageListArgs)
	realResult := result.(*chat.MessageServiceMessageListResult)
	success, err := handler.(chat.MessageService).MessageList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceMessageListArgs() interface{} {
	return chat.NewMessageServiceMessageListArgs()
}

func newMessageServiceMessageListResult() interface{} {
	return chat.NewMessageServiceMessageListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) MessagePost(ctx context.Context, req *chat.MessagePostRequest) (r *chat.MessagePostReponse, err error) {
	var _args chat.MessageServiceMessagePostArgs
	_args.Req = req
	var _result chat.MessageServiceMessagePostResult
	if err = p.c.Call(ctx, "MessagePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessageList(ctx context.Context, req *chat.MessageListRequest) (r *chat.MessageListResponse, err error) {
	var _args chat.MessageServiceMessageListArgs
	_args.Req = req
	var _result chat.MessageServiceMessageListResult
	if err = p.c.Call(ctx, "MessageList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
