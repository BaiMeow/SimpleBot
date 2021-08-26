//main.go暂时作为示例和调试使用
package main

import (
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/message"
	"log"
)

const addr = "ws://localhost:6700"
const token = ""

var b *bot.Bot

func main() {
	b = bot.New(driver.NewWsDriver(addr, token))
	b.Attach(&bot.GroupMsgHandler{
		Priority: 1,
		F:        justReply,
	})
	b.Attach(&bot.PrivateMsgHandler{
		Priority: 1,
		F:        justReply2,
	})
	b.Attach(&bot.GroupAddHandler{
		Priority: 1,
		F:        agree,
	})

	b.Run()
	select {}
}

//复读
func justReply(MsgID int32, GroupID int64, UserID int64, Msg message.Msg) bool {
	log.Println("new message")
	if msgid, err := b.SendGroupMsg(GroupID, Msg); err != nil {
		log.Println(err)
	} else {
		log.Println(msgid)
	}
	return false
}

//对纯文本消息，回复指定内容
func justReply2(MsgID int32, UserID int64, msg message.Msg) bool {
	if len(msg) != 1 || msg[0].GetType() != "text" {
		return false
	}
	log.Println("new message")
	if msgid, err := b.SendPrivateMsg(UserID, message.New().
		Text("BaiMeow").
		Image("https://baimeow.cn/me/BaiMeow.jpg", false).
		Face("1"),
	); err != nil {
		log.Println(err)
	} else {
		log.Println(msgid)
	}
	return false
}

func agree(request *bot.GroupRequest) bool {
	log.Println(request)
	request.Agree()
	return true
}
