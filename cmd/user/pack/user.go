package pack

import (
	"github.com/jiuxia211/ice-pomelo/cmd/user/dal/db"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
)

func User(data *db.User) (userResp *user.User) {
	if data == nil {
		return nil
	}

	return &user.User{
		Id:     data.ID,
		Name:   data.Username,
		Avatar: data.Avatar,
	}
}
