package bot

import (
	"encoding/json"
	"log"

	"github.com/BaiMeow/SimpleBot/handler"
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

	case "request":
	//不是Event那应该是api的回复
	case "":
		handleAPIReply(data, b)
	}
}

func handleGroupMsg(data []byte, b *Bot) {
	listener := b.listeners["message.group.normal"]
	if listener == nil {
		return
	}
	ev := new(groupEventFull)
	if err := handleUnmarshal(data, ev); err != nil {
		log.Println(err)
	}
	msg := ev.Message.ToMsgStruct()
	for _, v := range listener.heap {
		h, ok := v.(*handler.GroupMsgHandler)
		if !ok {
			continue
		}
		if h.F(ev.MessageID, ev.GroupID, ev.Sender.UserID, msg) {
			{
				return
			}
		}
	}
}

func handlePrivateMsg(data []byte, b *Bot) {
	listener := b.listeners["message.private.friend"]
	if listener == nil {
		return
	}
	ev := new(privateEventFull)
	if err := handleUnmarshal(data, ev); err != nil {
		log.Println(err)
	}
	msg := ev.Message.ToMsgStruct()
	for _, v := range listener.heap {
		h, ok := v.(*handler.PrivateMsgHandler)
		if !ok {
			continue
		}
		if h.F(ev.MessageID, ev.Sender.UserID, msg) {
			return
		}
	}
}
