package main

import (
	"context"
	"errors"

	"github.com/jiuxia211/ice-pomelo/cmd/user/pack"
	"github.com/jiuxia211/ice-pomelo/cmd/user/service"
	user "github.com/jiuxia211/ice-pomelo/kitex_gen/user"
	"github.com/jiuxia211/ice-pomelo/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	if _, err := utils.CheckApiToken(req.ApiToken); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errors.New("invalid user name or password"))
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Id = userResp.ID
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	if _, err := utils.CheckApiToken(req.ApiToken); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	token, err := utils.CreateToken(userResp.ID)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Token = token
	resp.Base = pack.BuildBaseResp(nil)
	resp.Id = userResp.ID
	return
}

// SendVerificationCode implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendVerificationCode(ctx context.Context, req *user.SendVerificationCodeRequest) (resp *user.SendVerificationCodeResponse, err error) {
	resp = new(user.SendVerificationCodeResponse)

	if _, err := utils.CheckApiToken(req.ApiToken); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	err = service.NewUserService(ctx).SendVertificationCode(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	resp = new(user.GetUserInfoResponse)

	if _, err := utils.CheckApiToken(req.ApiToken); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).GetUserInfo(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.User = pack.User(userResp)

	return
}
