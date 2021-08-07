package handler

type GroupMsgHandler struct {
	Priority int
	F        func(MsgID int32, GroupID, FromQQ int64, Msg string) bool
}

func (h *GroupMsgHandler) GetPriority() int {
	return h.Priority
}
