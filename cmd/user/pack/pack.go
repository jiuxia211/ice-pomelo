package pack

import (
	"github.com/jiuxia211/ice-pomelo/kitex_gen/user"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

func BuildBaseResp(err error) (resp *user.BaseResp) {
	resp = new(user.BaseResp)
	if err == nil {
		resp.Code = errz.SuccessCode
		resp.Msg = errz.SuccessMsg
	} else {
		resp.Code = errz.FailCode
		resp.Msg = err.Error()
	}
	return

}
