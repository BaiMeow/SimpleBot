package bot

import (
	"errors"
	"fmt"
	"github.com/BaiMeow/SimpleBot/driver"
	"log"
	"sort"
	"sync"
	"time"
)

type Bot struct {
	// qq号
	id     int64
	driver driver.Driver

	groupMsgListeners   *groupMsgHeap
	privateMsgListeners *privateMsgHeap
	listeners           map[string]*listenerHeap
}

type listener interface {
	getPriority() int
}

func New(d driver.Driver) *Bot {
	Bot := new(Bot)
	Bot.listeners = make(map[string]*listenerHeap)
	Bot.driver = d
	return Bot
}

func (b *Bot) Run() error {
	if err := b.driver.Run(); err != nil {
		return err
	}
	go func() {
		for {
			data, err := b.driver.Read()
			for err != nil {
				if errors.Is(err, driver.ErrConnClosed) {
					log.Println("bot was stopped")
					return
				}
				log.Println(err)
				time.Sleep(time.Second)
				err = b.driver.Run()
			}
			go handleEvent(data, b)
		}
	}()
	id, nickname, err := b.getLoginInfo()
	if err != nil {
		return err
	}
	b.id = id
	log.Println(fmt.Sprintf("已登陆到%s,qq号为%d", nickname, id))
	return nil
}

func (b *Bot) Attach(a listener) {
	var pos string
	// 单独处理群消息和私聊消息
	switch a.(type) {
	case *GroupMsgHandler:
		a := a.(*GroupMsgHandler)
		if b.groupMsgListeners == nil {
			b.groupMsgListeners = &groupMsgHeap{
				heap: []GroupMsgHandler{*a},
				lock: sync.Mutex{},
			}
			return
		}
		b.groupMsgListeners.lock.Lock()
		defer b.groupMsgListeners.lock.Unlock()
		b.groupMsgListeners.Push(a)
		sort.Sort(b.groupMsgListeners)
		return
	case *PrivateMsgHandler:
		a := a.(*PrivateMsgHandler)
		if b.privateMsgListeners == nil {
			b.privateMsgListeners = &privateMsgHeap{
				heap: []PrivateMsgHandler{*a},
				lock: sync.Mutex{},
			}
			return
		}
		b.privateMsgListeners.lock.Lock()
		defer b.privateMsgListeners.lock.Unlock()
		b.privateMsgListeners.Push(a)
		sort.Sort(b.privateMsgListeners)
		return
	case *GroupAddHandler:
		pos = "request.group.add"
	case *GroupInviteHandler:
		pos = "request.group.invite"
	case *GroupDecreaseHandler:
		pos = "notice.group_decrease"
	case *GroupKickMeHandler:
		pos = "notice.group_decrease.kick_me"
	case *GroupBanHandler:
		pos = "notice.group_ban.ban"
	case *GroupLiftBanHandler:
		pos = "notice.group_ban.lift_ban"
	}
	// 其他信息
	if b.listeners[pos] == nil {
		b.listeners[pos] = &listenerHeap{
			heap: []listener{a},
			lock: sync.Mutex{},
		}
		return
	}
	b.listeners[pos].lock.Lock()
	defer b.listeners[pos].lock.Unlock()
	b.listeners[pos].Push(&a)
	sort.Sort(b.listeners[pos])
}

// GetID 获取当前登陆的qq号
func (b *Bot) GetID() int64 {
	return b.id
}

func (b *Bot) Stop() error {
	return b.driver.Stop()
}
