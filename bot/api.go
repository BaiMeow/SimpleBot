package bot

import (
	"encoding/json"
	"errors"

	"github.com/BaiMeow/SimpleBot/message"
	"github.com/google/uuid"
)

//
var ErrJsonUnmarshal = errors.New("JsonUnmarshalError")

var waitReply = make(map[uuid.UUID]func([]byte, bool))

type preUnmarshalReply struct {
	Echo    uuid.UUID `json:"echo"`
	Retcode int       `json:"retcode"`
	Status  string    `json:"status"`
}

type apiCallFramework struct {
	Action  string      `json:"action"`
	Paramas interface{} `json:"params"`
	Echo    uuid.UUID   `json:"echo"`
}

type groupMsg struct {
	GroupID    int64       `json:"group_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}

type groupMsgReplyDetails struct {
	Data struct {
		MessageID int32 `json:"message_id"`
	} `json:"data"`
}

type privateMsg struct {
	UserID     int64       `json:"user_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}

type privateMsgReplyDetails struct {
	MessageID int32 `json:"message_id"`
}

func handleAPIReply(data []byte, b *Bot) {
	reply := new(preUnmarshalReply)
	if waitReply[reply.Echo] != nil {
		waitReply[reply.Echo](data, reply.Status == "ok")
		delete(waitReply, reply.Echo)
	}
}

//返回MsgID
func (b *Bot) SendGroupMsg(group int64, msg *message.Msg) (int32, error) {
	id := uuid.New()
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "send_group_msg",
		Paramas: groupMsg{
			GroupID:    group,
			Message:    msg.ToArrayMessage(),
			AutoEscape: false,
		},
		Echo: id,
	})
	if err != nil {
		return 0, err
	}
	b.driver.Write(bytes)
	msgid := make(chan int32, 1)
	waitReply[id] = func(data []byte, ok bool) {
		if !ok {
			msgid <- 0
			return
		}
		details := new(groupMsgReplyDetails)
		if err := json.Unmarshal(data, &details); err != nil {
			msgid <- 0
			return
		}
		msgid <- details.Data.MessageID
	}
	recMsgid := <-msgid
	if recMsgid == 0 {
		return 0, ErrJsonUnmarshal
	}
	return recMsgid, nil
}

func (b *Bot) SendPrivateMsg(qq int64, msg *message.Msg) (int32, error) {
	id := uuid.New()
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "send_private_msg",
		Paramas: privateMsg{
			UserID:     qq,
			Message:    msg.ToArrayMessage(),
			AutoEscape: false,
		},
		Echo: id,
	})
	if err != nil {
		return 0, err
	}
	b.driver.Write(bytes)
	msgid := make(chan int32, 1)
	waitReply[id] = func(data []byte, ok bool) {
		if !ok {
			msgid <- 0
		}
		details := new(privateMsgReplyDetails)
		if err := json.Unmarshal(data, details); err != nil {
			msgid <- 0
		}
		msgid <- details.MessageID
	}
	recMsgid := <-msgid
	if recMsgid == 0 {
		return 0, ErrJsonUnmarshal
	}
	return recMsgid, nil
}
