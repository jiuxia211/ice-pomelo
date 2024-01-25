package service

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"time"

	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/cache"
	"github.com/jiuxia211/ice-pomelo/config"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
	"gopkg.in/gomail.v2"
)

func generateRandomCode() string {
	// 生成6位数字验证码
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	return fmt.Sprintf("%d", code)
}

func (s *UserService) SendVertificationCode(req *user.SendVerificationCodeRequest) (err error) {
	dialer := gomail.NewDialer(config.ConfigInfo.Email.SmtpHost, 587, config.ConfigInfo.Email.SmtpEmail, config.ConfigInfo.Email.SmtpPass)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()

	verificationCode := generateRandomCode()
	//发送人
	m.SetHeader("From", config.ConfigInfo.Email.SmtpEmail)
	//接收人
	m.SetHeader("To", req.Email)
	//主题
	m.SetHeader("Subject", "[冰柚视频]|验证码")
	//内容
	m.SetBody("text/html", "[冰柚视频] 验证码:"+verificationCode+"\n用于注册冰柚视频账户，10分钟内有效。验证码泄露给他人可能导致账号被盗，请勿泄露，谨防被骗。")

	go cache.SetCode(s.ctx, req.Email, verificationCode)

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
