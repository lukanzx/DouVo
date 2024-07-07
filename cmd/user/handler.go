package main

import (
	"context"

	"github.com/lukanzx/DouVo/cmd/user/pack"
	"github.com/lukanzx/DouVo/cmd/user/service"
	"github.com/lukanzx/DouVo/kitex_gen/user"
	"github.com/lukanzx/DouVo/pkg/errno"
	"github.com/lukanzx/DouVo/pkg/utils"
)

type UserServiceImpl struct{}

func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)
	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	userResp, err := service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	token, err := utils.CreateToken(userResp.Id)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.UserId = userResp.Id
	resp.Token = token
	return
}

func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	userResp, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	token, err := utils.CreateToken(userResp.Id)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.User = pack.User(userResp)
	resp.Token = token
	return
}

func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	resp = new(user.InfoResponse)
	if req.UserId < 10000 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}
	userResp, err := service.NewUserService(ctx).GetUser(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.User = userResp
	return
}
