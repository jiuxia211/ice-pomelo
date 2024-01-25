package service

import (
	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/db"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
)

func (s *UserService) GetUserInfo(req *user.GetUserInfoRequest) (user *db.User, err error) {
	user, err = db.GetUser(s.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
