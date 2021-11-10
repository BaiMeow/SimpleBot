package bot

import (
	"encoding/json"
	"log"
)

type preUnmarshal struct {
	PostType      string `json:"post_type,omitempty"`
	MessageType   string `json:"message_type,omitempty"`
	NoticeType    string `json:"notice_type,omitempty"`
	RequestType   string `json:"request_type,omitempty"`
	MetaEventType string `json:"meta_event_type,omitempty"`
	SubType       string `json:"sub_type,omitempty"`
}

func handleEvent(data []byte, b *Bot) {
	preload := new(preUnmarshal)
	if err := json.Unmarshal(data, preload); err != nil {
		log.Println(err)
	}
	switch preload.PostType {
	case "meta_event":
		switch preload.MetaEventType {
		case "lifecycle":
			switch preload.SubType {
			case "connect":
				log.Println("连接成功")
			case "enable":
				log.Println("OneBot 启用")
			case "disable":
				log.Println("OneBot 停用")
			}
		case "heartbeat":
			//心跳
		}
	case "message":
		switch preload.MessageType {
		case "group":
			switch preload.SubType {
			case "normal":
				handleGroupMsg(data, b)
			}
		case "private":
			switch preload.SubType {
			case "friend":
				handlePrivateMsg(data, b)
			}

		}
	case "notice":
		switch preload.NoticeType {
		case "group_decrease":
			if preload.SubType == "kick_me" {
				handleGroupKickMe(data, b)
			} else if preload.SubType == "kick" || preload.SubType == "leave" {
				handleGroupDecrease(data, b)
			}
		case "group_ban":
			if preload.SubType == "ban" {
				handleGroupBan(data, b)
			} else if preload.SubType == "lift_ban" {
				handleGroupLiftBan(data, b)
			}
		}
	case "request":
		switch preload.RequestType {
		case "group":
			switch preload.SubType {
			case "add":
				handleGroupAdd(data, b)
			case "invite":
				handleGroupInvite(data, b)
			}
		case "friend":
			handleFriendAdd(data, b)
		}
	//不是Event那应该是api的回复
	case "":
		handleAPIReply(data)
	}
}

func handleGroupMsg(data []byte, b *Bot) {
	if b.groupMsgListeners == nil {
		return
	}
	ev := new(groupEventFull)
	if err := handleUnmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	msg := ev.Message.ToMsgStruct()
	for _, v := range b.groupMsgListeners.heap {
		if v.F(ev.MessageID, ev.GroupID, ev.Sender.UserID, msg) {

			return
		}
	}
}

func handlePrivateMsg(data []byte, b *Bot) {
	if b.privateMsgListeners == nil {
		return
	}
	ev := new(privateEventFull)
	if err := handleUnmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	msg := ev.Message.ToMsgStruct()
	for _, v := range b.privateMsgListeners.heap {
		if v.F(ev.MessageID, ev.Sender.UserID, msg) {
			return
		}
	}
}

func handleGroupAdd(data []byte, b *Bot) {
	if b.listeners["request.group.add"] == nil {
		return
	}
	listeners := b.listeners["request.group.add"]
	ev := new(groupReqEventFull)
	if err := json.Unmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	req := GroupRequest{
		handle:  b.respondGroupAdd,
		flag:    ev.Flag,
		UserID:  ev.UserID,
		GroupID: ev.GroupID,
		Comment: ev.Comment,
	}
	for _, v := range listeners.heap {
		v := v.(*GroupAddHandler)
		if v.F(&req) {
			return
		}
	}
}

func handleGroupInvite(data []byte, b *Bot) {
	if b.listeners["request.group.invite"] == nil {
		return
	}
	listeners := b.listeners["request.group.invite"]
	ev := new(groupReqEventFull)
	if err := json.Unmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	req := GroupRequest{
		handle:  b.respondGroupInvite,
		flag:    ev.Flag,
		UserID:  ev.UserID,
		GroupID: ev.GroupID,
		Comment: ev.Comment,
	}
	for _, v := range listeners.heap {
		v := v.(*GroupInviteHandler)
		if v.F(&req) {
			return
		}
	}
}

func handleGroupDecrease(data []byte, b *Bot) {
	if b.listeners["notice.group_decrease"] == nil {
		return
	}
	listener := b.listeners["notice.group_decrease"]
	ev := new(groupDecreaseFull)
	if err := json.Unmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	for _, v := range listener.heap {
		v := v.(*GroupDecreaseHandler)
		if v.F(ev.GroupID, ev.OperatorID, ev.UserID) {
			return
		}
	}
}

func handleGroupKickMe(data []byte, b *Bot) {
	if b.listeners["notice.group_decrease.kick_me"] == nil {
		return
	}
	listener := b.listeners["notice.group_decrease.kick_me"]
	ev := new(groupDecreaseFull)
	if err := json.Unmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	for _, v := range listener.heap {
		v := v.(*GroupKickMeHandler)
		if v.F(ev.GroupID, ev.OperatorID) {
			return
		}
	}
}

func handleGroupBan(data []byte, b *Bot) {
	if b.listeners["notice.group_ban.ban"] == nil {
		return
	}
	listener := b.listeners["notice.group_ban.ban"]
	ev := new(groupBanFull)
	if err := json.Unmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	for _, v := range listener.heap {
		v := v.(*GroupBanHandler)
		if v.F(ev.GroupID, ev.OperatorID, ev.UserID, ev.Duration) {
			return
		}
	}
}

func handleGroupLiftBan(data []byte, b *Bot) {
	if b.listeners["notice.group_ban.lift_ban"] == nil {
		return
	}
	listener := b.listeners["notice.group_ban.lift_ban"]
	ev := new(groupBanFull)
	if err := json.Unmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	for _, v := range listener.heap {
		v := v.(*GroupLiftBanHandler)
		if v.F(ev.GroupID, ev.OperatorID, ev.UserID) {
			return
		}
	}
}

func handleFriendAdd(data []byte, b *Bot) {
	if b.listeners["request.friend"] == nil {
		return
	}
	listeners := b.listeners["request.friend"]
	ev := new(friendAddFull)
	if err := json.Unmarshal(data, ev); err != nil {
		log.Println(err)
		return
	}
	req := FriendRequest{
		handle:  b.respondFriendAdd,
		flag:    ev.Flag,
		UserID:  ev.UserID,
		Comment: ev.Comment,
	}
	for _, v := range listeners.heap {
		v := v.(*FriendAddHandler)
		if v.F(&req) {
			return
		}
	}
}
