//main.go暂时作为示例和调试使用
package main

import (
	"log"

	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/handler"
	"github.com/BaiMeow/SimpleBot/message"
)

const addr = "ws://localhost:6700"
const token = ""

var b *bot.Bot

func main() {
	b = bot.New(driver.NewWsDriver(addr, token))
	b.Attach(&handler.GroupMsgHandler{
		Priority: 1,
		F:        justReply,
	})
	b.Attach(&handler.PrivateMsgHandler{
		Priority: 1,
		F:        justReply2,
	})
	b.Attach(&handler.GroupAddHandler{
		Priority: 1,
		F:        agree,
	})
	b.Run()
}

func justReply(MsgID int32, GroupID int64, UserID int64, Msg *message.Msg) bool {
	log.Println("new message")
	if msgid, err := b.SendGroupMsg(GroupID, Msg); err != nil {
		log.Println(err)
	} else {
		log.Panicln(msgid)
	}
	return false
}

func justReply2(MsgID int32, UserID int64, msg *message.Msg) bool {
	log.Println("new message")
	if msgid, err := b.SendPrivateMsg(UserID, msg); err != nil {
		log.Println(err)
	} else {
		log.Panicln(msgid)
	}
	return false
}

func agree(GroupID, UserID int64, comment, flag string) bool {
	log.Println(UserID)
	if err := b.RespondGroupAdd(true, flag, ""); err != nil {
		return false
	}
	return true
}
