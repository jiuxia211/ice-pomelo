package service

import (
	"errors"

	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/db"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CheckUser(req *user.LoginRequest) (user *db.User, err error) {
	user, err = db.GetUserByUsername(s.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, errors.New("密码错误")
	}

	return user, nil
}
