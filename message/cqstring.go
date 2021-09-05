package message

import (
	"strings"
)

func CQstrToArrayMessage(str string) *ArrayMessage {
	i := 0
	section := []string{""}
	for j, v := range str {
		switch v {
		case '[':
			if section[i] == "" {
				section[i] += "["
				continue
			}
			i++
			section = append(section, "[")
		case ']':
			section[i] += "]"
			if len(str)-1 != j {
				section = append(section, "")
				i++
			}
		default:
			section[i] += string(v)
		}

	}
	var msg ArrayMessage
	for _, v := range section {
		//不是cq码
		if len(v) < 6 || v[:4] != "[CQ:" {
			msg = append(msg, arrayMessageUnit{
				Type: "text",
				Data: map[string]string{
					"text": CQUnescape.Replace(v),
				},
			})
			continue
		}
		//处理cq码
		v = v[1 : len(v)-1]
		args := strings.Split(v, ",")
		var unit arrayMessageUnit
		unit.Type = args[0][3:]
		unit.Data = make(map[string]string)
		for _, v := range args {
			field := strings.SplitN(v, "=", 2)
			if len(field) != 2 {
				continue
			}
			unit.Data[CQUnescape.Replace(field[0])] = CQUnescape.Replace(field[1])
		}
		msg = append(msg, unit)
	}
	return &msg
}

var CQUnescape = strings.NewReplacer(
	"&amp;", "&",
	"&#91;", "[",
	"&#93;", "]",
)

var CQEscape = strings.NewReplacer(
	"&", "&amp;",
	"[", "&#91;",
	"]", "&#93;",
)
