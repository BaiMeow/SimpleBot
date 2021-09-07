package message

import "strconv"

func New() Msg {
	return Msg{}
}

// Text 纯文本消息
func (msg Msg) Text(txt string) Msg {
	return append(msg, Text{Text: txt})
}

// Image 在消息中添加图片;
//
// file 参数支持收到的文件的文件名,绝对路径(是onebot实例所在的计算机的绝对路径，不是本bot的,使用file URI格式),网络URL,Base64 编码;
//
// flash 闪照，传入true时发送闪照,闪照建议单独发送;
//
// 详细见 https://onebot.page.moe/v11/specs/message/segment.html#%E5%9B%BE%E7%89%87 ;
func (msg Msg) Image(file string, flash bool) Msg {
	return append(msg, Image{
		File: file,
		Type: func() string {
			if flash {
				return "flash"
			}
			return ""
		}(),
	})
}

// Face 发送qq表情,
// 表情id见 https://github.com/kyubotics/coolq-http-api/wiki/%E8%A1%A8%E6%83%85-CQ-%E7%A0%81-ID-%E8%A1%A8
func (msg Msg) Face(id string) Msg {
	return append(msg, Face{ID: id})
}

// At @成员,
// id为"all"时表示全体成员，支持传入int64与string
func (msg Msg) At(id interface{}) Msg {
	switch v := id.(type) {
	case string:
		if v == "all" {
			return append(msg, At{ID: "all"})
		}
	case int64:
		return append(msg, At{ID: strconv.FormatInt(v, 10)})
	}
	return msg
}

// Share 链接分享
// content与image为可选参数，建议单独发送该消息
func (msg Msg) Share(url, title, content, image string) Msg {
	return append(msg, Share{URL: url, Title: title, Content: content, Image: image})
}

// UserContact 推荐好友，建议单独发送该消息
func (msg Msg) UserContact(id int64) Msg {
	return append(msg, UserContact{ID: id})
}

//GroupContact 推荐群，建议单独发送该消息
func (msg Msg) GroupContact(id int64) Msg {
	return append(msg, GroupContact{ID: id})
}

//Location 位置，建议单独发送该消息
func (msg Msg) Location(lat, lon float64, title, content string) Msg {
	return append(msg, Location{Lat: lat, Lon: lon, Title: title, Content: content})
}

//Reply 回复，建议放在第一个
func (msg Msg) Reply(id int32) Msg {
	return append(msg, Reply{ID: id})
}
