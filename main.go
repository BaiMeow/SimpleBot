//main.go暂时作为示例和调试使用
package main

import (
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/handler"
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
	b.Run()
}

func justreply(MsgID int32, GroupID int64, FromQQ int64, Msg string) bool {
	b.SendGroupMsg(GroupID, Msg)
	return false
}
