// Code generated by hertz generator.

package api

import (
	"context"
	"io"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/jiuxia211/ice-pomelo/cmd/api/biz/model/api"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/pack"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/rpc"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
)

// Register .
// @router /pomelo/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.RegisterResponse)

	resp.ID, err = rpc.UserRegister(ctx, &user.RegisterRequest{
		Username:         req.Username,
		Password:         req.Password,
		Email:            req.Email,
		VerificationCode: req.VerificationCode,
	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}

// Login .
// @router /pomelo/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.LoginResponse)

	resp.ID, resp.Token, err = rpc.UserLogin(ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}

// SendVerificationCode .
// @router /pomelo/user/verification-code [POST]
func SendVerificationCode(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SendVerificationCodeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.SendVerificationCodeResponse)

	err = rpc.UserSendVerificationCode(ctx, &user.SendVerificationCodeRequest{
		Email: req.Email,
	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}

// GetUserInfo .
// @router /pomelo/user/get [POST]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetUserInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.GetUserInfoResponse)

	user, err := rpc.UserInfo(ctx, &user.GetUserInfoRequest{
		Token: req.Token,
		Id:    req.ID,
	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.User = pack.User(user)
	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}

// UploadUserAvatar .
// @router /pomelo/user/avatar [POST]
func UploadUserAvatar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UploadUserAvatarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		pack.SendFailResponse(c, err)
		return

	}
	// 简单的使用文件后缀来获取图片格式
	fileExt := filepath.Ext(file.Filename)
	fileContent, err := file.Open()
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	byteContainer, err := io.ReadAll(fileContent)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.UploadUserAvatarResponse)

	user, err := rpc.UserUploadAvatar(ctx, &user.UploadUserAvatarRequest{
		Token:  req.Token,
		Avatar: byteContainer,
		Format: fileExt,
	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.User = pack.User(user)
	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}
