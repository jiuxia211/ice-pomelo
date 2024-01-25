package service

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/db"
	"github.com/jiuxia211/ice-pomelo/config"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
	"github.com/jiuxia211/ice-pomelo/pkg/utils"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func (s *UserService) UploadUserAvater(req *user.UploadUserAvatarRequest, uid int64) (user *db.User, err error) {
	u, _ := url.Parse("https://" + config.ConfigInfo.Cos.BucketName + "." + config.ConfigInfo.Cos.Region)

	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("SECRETID"),
			SecretKey: os.Getenv("SECRETKEY"),
		},
	})

	key := fmt.Sprintf("%d_%s_avatar.png", uid, utils.GetTime())
	fileReader := bytes.NewReader(req.Avatar)
	avatar_url := "https://" + config.ConfigInfo.Cos.BucketName + "." + config.ConfigInfo.Cos.Region + "/" + key
	_, err = client.Object.Put(s.ctx, key, fileReader, nil)
	if err != nil {
		return nil, err
	}

	user, err = db.UpdateUserAvatar(s.ctx, uid, avatar_url)
	if err != nil {
		return nil, err
	}

	return user, err
}
