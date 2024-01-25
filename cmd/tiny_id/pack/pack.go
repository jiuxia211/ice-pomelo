package pack

import (
	"github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

func BuildBaseResp(err error) (resp *tiny_id.BaseResp) {
	resp = new(tiny_id.BaseResp)
	if err == nil {
		resp.Code = errz.SuccessCode
		resp.Msg = errz.SuccessMsg
	} else {
		resp.Code = errz.FailCode
		resp.Msg = err.Error()
	}
	return

}
