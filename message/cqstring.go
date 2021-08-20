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
		if v[0] != '[' {
			msg = append(msg, arrayMessageUnit{
				Type: "text",
				Data: map[string]interface{}{
					"text": replacer.Replace(v),
				},
			})
			continue
		}
		//处理cq码
		v = v[1 : len(v)-1]
		args := strings.Split(v, ",")
		var unit arrayMessageUnit
		unit.Type = args[0][3:]
		unit.Data = make(map[string]interface{})
		for _, v := range args {
			field := strings.SplitN(v, "=", 2)
			if len(field) != 2 {
				continue
			}
			unit.Data[replacer.Replace(field[0])] = replacer.Replace(field[1])
		}
		msg = append(msg, unit)
	}
	return &msg
}

var replacer = strings.NewReplacer(
	"&amp;", "&",
	"&#91;", "[",
	"&#93;", "]",
)
