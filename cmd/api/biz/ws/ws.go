package ws

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/ws/db"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/ws/model"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/ws/pack"
	"github.com/jiuxia211/ice-pomelo/pkg/errz"
)

type Client struct {
	ID   int64
	Conn *websocket.Conn
}

// 我们维护一个manager，它需要保存所有连接的用户信息，以及所有操作的信号
type WsManager struct {
	Clients    map[int64]*Client   // 连接的用户map
	Register   chan *Client        // 用户注册信号
	Unregister chan *Client        // 用户退出信号
	Send       chan *model.SendMsg // 发消息信号
}

var Manager = WsManager{
	Clients:    make(map[int64]*Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Send:       make(chan *model.SendMsg),
}

func (wsManager *WsManager) Start() {

	for {
		select {
		case signal := <-wsManager.Register:
			wsManager.Clients[signal.ID] = signal

		case signal := <-wsManager.Unregister:
			if _, ok := wsManager.Clients[signal.ID]; ok {
				reply := &model.ReplyMsg{
					Code:    errz.SuccessCode,
					Content: "连接中断",
					From:    0,
				}
				signal.Conn.WriteJSON(reply)
			}
		case signal := <-wsManager.Send:
			msg := &db.Message{
				UID:     signal.UID,
				ToUID:   signal.ToUID,
				Content: signal.Content,
			}
			if _, ok := wsManager.Clients[signal.ToUID]; ok {
				// 对方在线
				reply := &model.ReplyMsg{
					Code:    errz.SuccessCode,
					Content: signal.Content,
					From:    signal.UID,
				}
				wsManager.Clients[signal.ToUID].Conn.WriteJSON(&reply)
				msg.Read = true

			} else {
				// 对方不在线
				msg.Read = false
			}

			err := db.CreateMessage(msg)
			if err != nil {
				hlog.Debugf("存储消息失败")
				wsManager.Clients[signal.UID].Conn.WriteJSON(pack.BuildReplyMsg(err))
			}
		}
	}
}
func (c *Client) Read(ctx context.Context, uid int64) {
	defer func() {
		Manager.Unregister <- c
		c.Conn.Close()
	}()
	for {
		sendMsg := new(model.SendMsg)
		err := c.Conn.ReadJSON(&sendMsg)
		if err != nil {
			hlog.Debugf("用户%v已下线\n", c.ID)
			break
		}
		switch sendMsg.Type {
		case 1: // 私聊
			// TODO 防止刷消息
			sendMsg.UID = uid
			Manager.Send <- sendMsg
		case 2: // 获取历史记录
			msgList, err := db.GetAllMessage(ctx, uid, sendMsg.ToUID)
			if err != nil {
				c.Conn.WriteJSON(pack.BuildReplyMsg(err))
			} else {
				c.Conn.WriteJSON(msgList)
			}
		case 3: // 获取未读记录
			msgList, err := db.GetUnreadMessage(ctx, uid)
			if err != nil {
				c.Conn.WriteJSON(pack.BuildReplyMsg(err))
			} else {
				c.Conn.WriteJSON(msgList)
			}
		}
	}
}
