package pack

import (
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/ws/model"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

func BuildReplyMsg(err error) (reply *model.ReplyMsg) {
	reply = new(model.ReplyMsg)
	if err == nil {
		reply.Code = errz.SuccessCode
		reply.Content = errz.SuccessMsg
	} else {
		reply.Code = errz.FailCode
		reply.Content = err.Error()
	}
	return
}
