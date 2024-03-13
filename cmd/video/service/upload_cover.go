package service

import (
	"bytes"
	"net/http"
	"net/url"
	"os"

	"github.com/jiuxia211/ice-pomelo/config"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/video"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func (s *VideoService) UploadCover(req *video.UploadVideoRequest, key string) (err error) {
	u, _ := url.Parse("https://" + config.ConfigInfo.Cos.BucketName + "." + config.ConfigInfo.Cos.Region)

	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("SECRETID"),
			SecretKey: os.Getenv("SECRETKEY"),
		},
	})

	fileReader := bytes.NewReader(req.CoverFile)

	_, err = client.Object.Put(s.ctx, key, fileReader, nil)
	if err != nil {
		return err
	}

	return nil
}
