package bot

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

var waitReply = make(map[uuid.UUID]func([]byte, bool))

type preUnmarshalReply struct {
	Echo    uuid.UUID `json:"echo"`
	Retcode int       `json:"retcode"`
	Status  string    `json:"status"`
}

type groupMsgReplyDetails struct {
	Data struct {
		MsgID int32 `json:"message_id"`
	} `json:"data"`
}

func handleAPIReply(data []byte, b *Bot) {
	reply := new(preUnmarshalReply)
	fmt.Println(json.Unmarshal(data, reply))
	fmt.Println(reply)
	if waitReply[reply.Echo] != nil {
		waitReply[reply.Echo](data, reply.Status == "ok")
		delete(waitReply, reply.Echo)
	}
}

//返回MsgID
func (b *Bot) SendGroupMsg(group int64, msg string) (int32, error) {
	id := uuid.New()
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "send_group_msg",
		Paramas: groupMsg{
			GroupID:    group,
			Message:    msg,
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
		msgid <- details.Data.MsgID
	}
	return <-msgid, nil
}

type apiCallFramework struct {
	Action  string      `json:"action"`
	Paramas interface{} `json:"params"`
	Echo    uuid.UUID   `json:"echo"`
}

type groupMsg struct {
	GroupID    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
