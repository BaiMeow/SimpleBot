package handler

import "github.com/BaiMeow/SimpleBot/message"

type GroupMsgHandler struct {
	Priority int
	F        func(MsgID int32, GroupID, FromQQ int64, Msg *message.Msg) bool
}

func (h *GroupMsgHandler) GetPriority() int {
	return h.Priority
}

type PrivateMsgHandler struct {
	Priority int
	F        func(MsgID int32, FromQQ int64, Msg *message.Msg) bool
}

func (h *PrivateMsgHandler) GetPriority() int {
	return h.Priority
}
