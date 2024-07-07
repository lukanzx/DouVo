package pack

import (
	"errors"

	"github.com/lukanzx/DouVo/kitex_gen/chat"
	"github.com/lukanzx/DouVo/pkg/errno"
)

func BuildBaseResp(err error) *chat.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *chat.BaseResp {
	return &chat.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
