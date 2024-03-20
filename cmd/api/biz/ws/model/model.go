package model

type SendMsg struct {
	Type    int64 // type1 私聊 type2 获取和某个用户的所有历史消息 type3获取所有(不同用户)未读消息
	UID     int64 // 这个变量由token获取
	ToUID   int64
	Content string
}
type ReplyMsg struct {
	Code    int64
	Content string
	From    int64
}
