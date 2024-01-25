package pack

import (
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/model/api"
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
)

func User(data *user.User) (userResp *api.User) {
	if data == nil {
		return nil
	}

	return &api.User{
		ID:        data.Id,
		Name:      data.Name,
		AvatarURL: data.AvatarUrl,
	}
}
