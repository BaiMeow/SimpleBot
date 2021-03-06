package bot

import (
	"encoding/json"
	"errors"
	"github.com/BaiMeow/SimpleBot/message"
	"github.com/google/uuid"
	"log"
)

// ErrJsonUnmarshal json序列化中出错
var ErrJsonUnmarshal = errors.New("JsonUnmarshalError")

var waitReply = make(map[string]func([]byte, bool))

type preUnmarshalReply struct {
	Echo    string `json:"echo"`
	RetCode int    `json:"retcode"`
	Status  string `json:"status"`
}

type apiCallFramework struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
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
	Data struct {
		MessageID int32 `json:"message_id"`
	} `json:"data"`
}

type groupAddOrInvite struct {
	Flag    string `json:"flag"`
	SubType string `json:"sub_type"`
	Approve bool   `json:"approve"`
	Reason  string `json:"reason"`
}

type friendAdd struct {
	Flag    string `json:"flag"`
	Approve bool   `json:"approve"`
	Remark  string `json:"remark"`
}

type loginInfoDetails struct {
	Data struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"data"`
}

type groupBan struct {
	GroupID  int64 `json:"group_id"`
	UserID   int64 `json:"user_id"`
	Duration int64 `json:"duration"`
}

func handleAPIReply(data []byte) {
	reply := new(preUnmarshalReply)
	if err := json.Unmarshal(data, reply); err != nil {
		log.Println(err)
	}
	if reply.Echo == "" {
		log.Println("未知的api响应")
		return
	}
	if reply.Status != "ok" {
		log.Println("onebot服务处理失败")
		if waitReply[reply.Echo] != nil {
			waitReply[reply.Echo](data, false)
			delete(waitReply, reply.Echo)
		}
		return
	}
	if waitReply[reply.Echo] != nil {
		waitReply[reply.Echo](data, true)
		delete(waitReply, reply.Echo)
	}
}

// SendGroupMsg 发送群聊消息(不含匿名消息)
func (b *Bot) SendGroupMsg(group int64, msg message.Msg) (int32, error) {
	id := uuid.New().String()
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "send_group_msg",
		Params: groupMsg{
			GroupID:    group,
			Message:    msg.ToArrayMessage(),
			AutoEscape: false,
		},
		Echo: id,
	})
	if err != nil {
		return 0, err
	}

	msgID := make(chan int32, 1)
	waitReply[id] = func(data []byte, ok bool) {
		if !ok {
			msgID <- 0
			return
		}
		details := new(groupMsgReplyDetails)
		if err := json.Unmarshal(data, &details); err != nil {
			log.Println(err)
			msgID <- 0
			return
		}
		msgID <- details.Data.MessageID
	}
	if err := b.writeWithRetry(bytes); err != nil {
		delete(waitReply, id)
		return 0, err
	}
	recMsgID := <-msgID
	if recMsgID == 0 {
		return 0, ErrJsonUnmarshal
	}
	return recMsgID, nil
}

// SendPrivateMsg 发送私聊消息
func (b *Bot) SendPrivateMsg(qq int64, msg message.Msg) (int32, error) {
	id := uuid.New().String()
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "send_private_msg",
		Params: privateMsg{
			UserID:     qq,
			Message:    msg.ToArrayMessage(),
			AutoEscape: false,
		},
		Echo: id,
	})
	if err != nil {
		return 0, err
	}
	msgID := make(chan int32, 1)
	waitReply[id] = func(data []byte, ok bool) {
		if !ok {
			msgID <- 0
			return
		}
		details := new(privateMsgReplyDetails)
		if err := json.Unmarshal(data, details); err != nil {
			log.Println(err)
			msgID <- 0
			return
		}
		msgID <- details.Data.MessageID
	}
	if err := b.writeWithRetry(bytes); err != nil {
		delete(waitReply, id)
		return 0, err
	}
	recMsgID := <-msgID
	if recMsgID == 0 {
		return 0, ErrJsonUnmarshal
	}
	return recMsgID, nil
}

func (b *Bot) respondGroupAdd(approve bool, flag, reason string) error {
	if approve {
		reason = ""
	}
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "set_group_add_request",
		Params: groupAddOrInvite{
			Flag:    flag,
			SubType: "add",
			Approve: approve,
			Reason:  reason,
		},
	})
	if err != nil {
		return err
	}
	return b.writeWithRetry(bytes)
}

func (b *Bot) respondGroupInvite(approve bool, flag, reason string) error {
	if approve {
		reason = ""
	}
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "set_group_add_request",
		Params: groupAddOrInvite{
			Flag:    flag,
			SubType: "invite",
			Approve: approve,
			Reason:  reason,
		},
	})
	if err != nil {
		return err
	}
	return b.writeWithRetry(bytes)
}

func (b *Bot) respondFriendAdd(approve bool, flag, remark string) error {
	if !approve {
		remark = ""
	}
	id := uuid.NewString()
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "set_friend_add_request",
		Params: friendAdd{
			Flag:    flag,
			Approve: approve,
			Remark:  remark,
		},
		Echo: id,
	})
	if err != nil {
		return err
	}
	ok := make(chan bool)
	waitReply[id] = func(bytes []byte, b bool) {
		ok <- b
	}
	if err := b.writeWithRetry(bytes); err != nil {
		delete(waitReply, id)
		return err
	}
	if !<-ok {
		return errors.New("fail to handle friend add request")
	}
	return nil
}

func (b *Bot) getLoginInfo() (int64, string, error) {
	id := uuid.New().String()
	bytes, err := json.Marshal(
		&apiCallFramework{
			Action: "get_login_info",
			Params: nil,
			Echo:   id,
		})
	if err != nil {
		return 0, "", err
	}
	info := make(chan *loginInfoDetails, 1)
	waitReply[id] = func(data []byte, ok bool) {
		if !ok {
			info <- nil
			return
		}
		i := new(loginInfoDetails)
		if err := json.Unmarshal(data, i); err != nil {
			log.Println(err)
			info <- nil
			return
		}
		info <- i
	}
	if err := b.writeWithRetry(bytes); err != nil {
		delete(waitReply, id)
		return 0, "", err
	}
	recInfo := <-info
	if recInfo != nil {
		return recInfo.Data.UserID, recInfo.Data.Nickname, nil
	}
	return 0, "", ErrJsonUnmarshal
}

//SetGroupBan 群禁言，duration单位为秒，传入0时解除禁言
//
//但是QQ貌似是按照分钟计算的，因此建议传入60的倍数
func (b *Bot) SetGroupBan(GroupID, UserID, Duration int64) error {
	id := uuid.New().String()
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "set_group_ban",
		Params: groupBan{
			GroupID:  GroupID,
			UserID:   UserID,
			Duration: Duration,
		},
		Echo: id,
	})
	if err != nil {
		return err
	}
	return b.writeWithRetry(bytes)
}
