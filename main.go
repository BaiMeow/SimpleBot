//main.go暂时作为示例和调试使用
package main

import (
	"log"

	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/handler"
	"github.com/BaiMeow/SimpleBot/message"
)

var addr = "ws://localhost:6700"
var token = ""

var b *bot.Bot

func main() {
	b = bot.New(driver.NewWsDriver(addr, token))
	b.Attach("message.group.normal", &handler.GroupMsgHandler{
		Priority: 1,
		F:        justreply,
	})
	b.Attach("message.private.friend", &handler.PrivateMsgHandler{
		Priority: 1,
		F:        justreply2,
	})
	b.Run()
}

func justreply(MsgID int32, GroupID int64, FromQQ int64, Msg *message.Msg) bool {
	log.Println("new message")
	if msgid, err := b.SendGroupMsg(GroupID, Msg); err != nil {
		log.Println(err)
	} else {
		log.Panicln(msgid)
	}
	return false
}

func justreply2(msgid int32, fromqq int64, msg *message.Msg) bool {
	log.Println("new message")
	if msgid, err := b.SendPrivateMsg(fromqq, msg); err != nil {
		log.Println(err)
	} else {
		log.Panicln(msgid)
	}
	return false
}
