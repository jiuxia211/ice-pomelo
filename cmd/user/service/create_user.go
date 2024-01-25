package service

import (
	"errors"

	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/cache"
	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/db"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *UserService) CreateUser(req *user.RegisterRequest) (userResp *db.User, err error) {

	// 1.验证码校验
	vertified, err := cache.CheckCode(s.ctx, req.Email, req.VerificationCode)
	if err == redis.Nil {
		return nil, errors.New("不存在邮箱对应的验证码")
	}
	if err != nil {
		return nil, err
	}
	if !vertified {
		return nil, errors.New("验证码错误")
	}
	// 2.密码加密
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// 3.确认用户名是否唯一
	_, err = db.GetUserByUsername(s.ctx, req.Username)
	if err == gorm.ErrRecordNotFound {
		// 4.获取可用id
		id, err := getID(s.ctx, req.ApiToken)
		if err != nil {
			return nil, err
		}
		userModel := &db.User{
			ID:       id,
			Username: req.Username,
			Password: string(hashBytes),
			Email:    req.Email,
			Avatar:   "todo",
		}

		userResp, err = db.CreateUser(s.ctx, userModel)
		if err != nil {
			return nil, err
		}

		return userResp, nil
	} else if err != nil {
		return nil, err
	} else {
		err = errors.New("用户已经存在")
		return nil, err
	}

}
