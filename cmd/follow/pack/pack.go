package pack

import (
	"errors"

	"github.com/lukanzx/DouVo/kitex_gen/follow"
	"github.com/lukanzx/DouVo/pkg/errno"
)

func BuildBaseResp(err error) *follow.BaseResp {
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

func baseResp(err errno.ErrNo) *follow.BaseResp {
	return &follow.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
