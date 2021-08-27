package bot

import "github.com/BaiMeow/SimpleBot/message"

//虽然mirai传了一堆参数，但是用得到的毕竟是少数
type groupEventFull struct {
	Time      int64 `json:"time"`
	SelfID    int64 `json:"self_id"`
	MessageID int32 `json:"message_id"`
	GroupID   int64 `json:"group_id"`
	UserID    int64 `json:"user_id"`
	Anonymous struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	} `json:"anonymous,omitempty"`
	Message    message.ArrayMessage `json:"message"`
	RawMessage string               `json:"raw_message"`
	Font       int32                `json:"font"`
	Sender     struct {
		UserID   int64  `json:"user_id"`
		NickName string `json:"nickname"`
		Sex      string `json:"sex"`
		Age      int32  `json:"age"`
	}
}

type privateEventFull struct {
	Time       int64                `json:"time"`
	SelfID     int64                `json:"self_id"`
	MessageID  int32                `json:"message_id"`
	UserID     int64                `json:"user_id"`
	Message    message.ArrayMessage `json:"message"`
	RawMessage string               `json:"raw_message"`
	Font       int32                `json:"font"`
	Sender     struct {
		UserID   int64  `json:"user_id"`
		NickName string `json:"nickname"`
		Sex      string `json:"sex"`
		Age      int32  `json:"age"`
	} `json:"sender"`
}

type groupReqEventFull struct {
	Time    int64  `json:"time"`
	SelfID  int64  `json:"self_id"`
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
}

type groupDecreaseFull struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
	SubType    string `json:"sub_type"`
}

func (f *groupEventFull) getMessage() *message.ArrayMessage {
	return &f.Message
}

func (f *privateEventFull) getMessage() *message.ArrayMessage {
	return &f.Message
}
