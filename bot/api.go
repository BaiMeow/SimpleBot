package bot

import (
	"encoding/json"
)

func (b *Bot) SendGroupMsg(group int64, msg string) error {
	bytes, err := json.Marshal(&apiCallFramework{
		Action: "send_group_msg",
		Paramas: groupMsg{
			GroupID:    group,
			Message:    msg,
			AutoEscape: false,
		},
		Echo: "",
	})
	if err != nil {
		return err
	}
	b.driver.Write(bytes)
	return nil

}

type apiCallFramework struct {
	Action  string      `json:"action"`
	Paramas interface{} `json:"params"`
	Echo    string      `json:"echo"`
}

type groupMsg struct {
	GroupID    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
