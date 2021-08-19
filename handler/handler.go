package handler

import "github.com/BaiMeow/SimpleBot/message"

//GroupMsgHandler 接收处理群聊消息（不含匿名消息）
type GroupMsgHandler struct {
	Priority int
	F        func(MsgID int32, GroupID, FromQQ int64, Msg *message.Msg) bool
}

func (h *GroupMsgHandler) GetPriority() int {
	return h.Priority
}

//PrivateMsgHandler 接收处理私聊消息
type PrivateMsgHandler struct {
	Priority int
	F        func(MsgID int32, FromQQ int64, Msg *message.Msg) bool
}

func (h *PrivateMsgHandler) GetPriority() int {
	return h.Priority
}
