package handler

import "github.com/BaiMeow/SimpleBot/message"

//GroupMsgHandler 接收处理群聊消息（不含匿名消息）
type GroupMsgHandler struct {
	Priority int
	F        func(MsgID int32, GroupID, UserID int64, Msg message.Msg) bool
}

//PrivateMsgHandler 接收处理私聊消息
type PrivateMsgHandler struct {
	Priority int
	F        func(MsgID int32, UserID int64, Msg message.Msg) bool
}

//GroupAddHandler 处理加群申请
type GroupAddHandler struct {
	Priority int
	F        func(GroupID, UserID int64, comment, flag string) bool
}

func (h *GroupMsgHandler) GetPriority() int {
	return h.Priority
}

func (h *PrivateMsgHandler) GetPriority() int {
	return h.Priority
}

func (h *GroupAddHandler) GetPriority() int {
	return h.Priority
}
