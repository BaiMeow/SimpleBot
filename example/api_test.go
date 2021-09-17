package main

import (
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"testing"
	"time"
)

func TestGroupBan(t *testing.T) {
	b := bot.New(driver.NewWsDriver(addr, token))
	b.Run()
	b.SetGroupBan(452384598, 1098105012, 1)
	time.Sleep(5)
	b.SetGroupBan(452384598, 1098105012, 0)
}
