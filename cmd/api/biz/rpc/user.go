package rpc

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user/userservice"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

func InitUserRPC() {
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithHostPorts("127.0.0.1:8211"))
	if err != nil {
		klog.Fatalf("user init rpc error: %v", err)
	}
	userClient = c
}

func UserRegister(ctx context.Context, req *user.RegisterRequest) (int64, error) {
	resp, err := userClient.Register(ctx, req)

	if err != nil {
		return -1, err
	}

	if resp.Base.Code != errz.SuccessCode {
		return 0, errors.New(resp.Base.Msg)
	}

	return resp.Id, nil
}

func UserLogin(ctx context.Context, req *user.LoginRequest) (int64, string, error) {
	resp, err := userClient.Login(ctx, req)

	if err != nil {
		return -1, "", err
	}

	if resp.Base.Code != errz.SuccessCode {
		return 0, "", errors.New(resp.Base.Msg)
	}

	return resp.Id, resp.Token, nil
}

func UserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.User, error) {
	resp, err := userClient.GetUserInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errz.SuccessCode {
		return nil, errors.New(resp.Base.Msg)
	}

	return resp.User, nil
}

func UserSendVerificationCode(ctx context.Context, req *user.SendVerificationCodeRequest) error {
	resp, err := userClient.SendVerificationCode(ctx, req)

	if err != nil {
		return err
	}
	if resp.Base.Code != errz.SuccessCode {
		return errors.New(resp.Base.Msg)
	}

	return nil
}
