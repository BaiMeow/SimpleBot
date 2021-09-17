package main

import (
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"testing"
)

func TestGroupAdd(t *testing.T) {
	b := bot.New(driver.NewWsDriver(addr, token))
	b.Attach(&bot.GroupAddHandler{
		Priority: 1,
		F: func(request *bot.GroupRequest) bool {
			t.Log("加群消息")
			return false
		},
	})
	b.Run()
	select {}
}
