package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/model/api"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

type Response struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func BuildBaseResp(err error) (resp *api.BaseResp) {
	resp = new(api.BaseResp)
	if err == nil {
		resp.Code = errz.SuccessCode
		resp.Msg = errz.SuccessMsg
	} else {
		resp.Code = errz.FailCode
		resp.Msg = err.Error()
	}
	return

}

func SendFailResponse(c *app.RequestContext, err error) {

	c.JSON(consts.StatusOK, Response{
		Code: errz.FailCode,
		Msg:  err.Error(),
	})
}

func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}
