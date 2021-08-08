package bot

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MscBaiMeow/SimpleBot/handler"
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
			handleGroupEvent(data, b)
		case "private":
			handlePrivateEvent(data, b)
		}
	case "notice":

	case "request":
	//不是Event那应该是api的回复
	case "":
		handleAPIReply(data, b)
	}
}

//虽然mirai传了一堆参数，但是用得到的毕竟是少数
type groupEventFull struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageID   int32  `json:"message_id"`
	GroupID     int64  `json:"group_id"`
	UserID      int64  `json:"user_id"`
	Anonymous   struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	} `json:"anonymous,omitempty"`
	Message    string
	RawMessage string `json:"raw_message"`
	Font       int32  `json:"font"`
	Sender     struct {
		UserID   int64  `json:"user_id"`
		NickName string `json:"nickname"`
		Sex      string `json:"sex"`
		Age      int32  `json:"age"`
	}
}

func handleGroupEvent(data []byte, b *Bot) {
	ev := new(groupEventFull)
	json.Unmarshal(data, ev)
	list := b.listeners[fmt.Sprintf("%s.%s.%s", ev.PostType, ev.MessageType, ev.SubType)]
	for _, v := range list.heap {
		h, ok := v.(*handler.GroupMsgHandler)
		if !ok {
			continue
		}
		if h.F(ev.MessageID, ev.GroupID, ev.Sender.UserID, ev.Message) {
			return
		}
	}
}

func handlePrivateEvent(data []byte, b *Bot) {}
