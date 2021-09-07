package main

import (
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/message"
	"testing"
)

func TestMsg_GroupContact(t *testing.T) {
	b := bot.New(driver.NewWsDriver(addr, token))
	b.Run()
	_, err := b.SendPrivateMsg(1098105012, message.New().GroupContact(609632487))
	if err != nil {
		t.Error(err)
	}
}

func TestMsg_Share(t *testing.T) {
	b := bot.New(driver.NewWsDriver(addr, token))
	b.Run()
	_, err := b.SendPrivateMsg(1098105012, message.New().Share("https://www.bilibili.com/", "", "", ""))
	if err != nil {
		t.Error(err)
	}
}

func TestMsg_Location(t *testing.T) {
	b := bot.New(driver.NewWsDriver(addr, token))
	b.Run()
	_, err := b.SendGroupMsg(1098105012, message.New().Location(120, 30, "杭州", "城市"))
	if err != nil {
		t.Error(err)
	}
}

func TestMsg_Reply(t *testing.T) {
	b := bot.New(driver.NewWsDriver(addr, token))
	b.Attach(&bot.GroupMsgHandler{
		Priority: 1,
		F: func(MsgID int32, GroupID, UserID int64, Msg message.Msg) bool {
			_, err := b.SendGroupMsg(GroupID, message.New().Reply(MsgID).Text("回复测试"))
			if err != nil {
				t.Error(err)
			}
			return true
		},
	})
	b.Run()
	select {}
}
