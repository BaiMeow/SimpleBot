package message

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
// flash 闪照，传入true时发送闪照,尽量不要和其他类型消息一起发送;
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
// id为"all"时表示全体成员(所以用了string类型)
func (msg Msg) At(id string) Msg {
	return append(msg, At{ID: id})
}
