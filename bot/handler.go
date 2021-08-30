package bot

import (
	"github.com/BaiMeow/SimpleBot/message"
)

// GroupMsgHandler 群聊消息（不含匿名消息）
type GroupMsgHandler struct {
	Priority int
	F        func(MsgID int32, GroupID, UserID int64, Msg message.Msg) bool
}

// PrivateMsgHandler 私聊消息
type PrivateMsgHandler struct {
	Priority int
	F        func(MsgID int32, UserID int64, Msg message.Msg) bool
}

// GroupAddHandler 加群申请
type GroupAddHandler struct {
	Priority int
	F        func(request *GroupRequest) bool
}

// GroupInviteHandler 加群邀请
type GroupInviteHandler struct {
	Priority int
	F        func(request *GroupRequest) bool
}

// GroupDecreaseHandler 群成员减少,包含成员主动退群，成员被踢;
//
// 主动退群时OperatorID==UserID，不相等则为被踢
type GroupDecreaseHandler struct {
	Priority int
	F        func(GroupID, OperatorID, UserID int64) bool
}

// GroupKickMeHandler 自己被踢出群聊
type GroupKickMeHandler struct {
	Priority int
	F        func(GroupID, OperatorID int64) bool
}

func (h *GroupMsgHandler) getPriority() int {
	return h.Priority
}

func (h *PrivateMsgHandler) getPriority() int {
	return h.Priority
}

func (h *GroupAddHandler) getPriority() int {
	return h.Priority
}

func (h *GroupInviteHandler) getPriority() int {
	return h.Priority
}

func (h *GroupDecreaseHandler) getPriority() int {
	return h.Priority
}

func (h *GroupKickMeHandler) getPriority() int {
	return h.Priority
}
