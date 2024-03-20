package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/pack"
	"github.com/jiuxia211/ice-pomelo/cmd/api/biz/ws"
	"github.com/jiuxia211/ice-pomelo/pkg/utils"
)

var upgrader = websocket.HertzUpgrader{}

func WebsocketHandler(ctx context.Context, c *app.RequestContext) {
	token := c.GetHeader("Authorization")

	claim, err := utils.CheckToken(string(token))
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	err = upgrader.Upgrade(c, func(conn *websocket.Conn) {
		client := &ws.Client{
			ID:   claim.UserId,
			Conn: conn,
		}
		ws.Manager.Register <- client
		client.Read(ctx, claim.UserId)

	})
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

}
