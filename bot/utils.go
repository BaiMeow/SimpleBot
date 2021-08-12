package bot

import (
	"encoding/json"
	"strings"

	"github.com/BaiMeow/SimpleBot/message"
)

type eventFull interface {
	getMessage() *message.ArrayMessage
}

func suckCQstr(data []byte) (cqstr string) {
	str := string(data)
	i := strings.Index(str, "\"message\":") + 11
	j := i
	for ; !(str[j] == '"' && str[j-1] != '\\') && j < len(str); j++ {
	}
	if j == len(str) {
		return ""
	}
	return str[i:j]
}

func handleUnmarshal(data []byte, ev eventFull) error {
	if err := json.Unmarshal(data, ev); err != nil {
		if errForm, ok := err.(*json.UnmarshalTypeError); ok && errForm.Field == "message" && errForm.Value == "string" {
			*ev.getMessage() = *message.CQstrToArrayMessage(suckCQstr(data))
			return nil
		}
		return err
	}
	return nil
}
