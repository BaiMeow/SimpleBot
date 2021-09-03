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
