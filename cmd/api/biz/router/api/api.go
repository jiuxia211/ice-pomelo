// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "github.com/jiuxia211/ice-pomelo/cmd/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_pomelo := root.Group("/pomelo", _pomeloMw()...)
		_pomelo.PUT("/video", append(_uploadvideoMw(), api.UploadVideo)...)
		_video := _pomelo.Group("/video", _videoMw()...)
		_video.GET("/feed", append(_feedMw(), api.Feed)...)
		_pomelo.GET("/user", append(_getuserinfoMw(), api.GetUserInfo)...)
		_user := _pomelo.Group("/user", _userMw()...)
		_user.PUT("/avatar", append(_uploaduseravatarMw(), api.UploadUserAvatar)...)
		{
			_user0 := _pomelo.Group("/user", _user0Mw()...)
			_user0.POST("/login", append(_loginMw(), api.Login)...)
			_user0.POST("/register", append(_registerMw(), api.Register)...)
			_user0.POST("/verification-code", append(_sendverificationcodeMw(), api.SendVerificationCode)...)
		}
	}
}
