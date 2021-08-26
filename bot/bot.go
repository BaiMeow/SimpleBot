package bot

import (
	"github.com/BaiMeow/SimpleBot/driver"
	"sort"
)

type Bot struct {
	driver driver.Driver

	groupMsgListeners   *groupMsgHeap
	privateMsgListeners *privateMsgHeap
	listeners           map[string]*listenerHeap
}

type listener interface {
	GetPriority() int
}

func New(d driver.Driver) *Bot {
	Bot := new(Bot)
	Bot.listeners = make(map[string]*listenerHeap)
	Bot.driver = d
	return Bot
}

func (b *Bot) Run() {
	b.driver.Run()
	for {
		go handleEvent(b.driver.Read(), b)
	}
}

func (b *Bot) Attach(a listener) {
	var pos string
	//单独处理群消息和私聊消息
	switch a.(type) {
	case *GroupMsgHandler:
		a := a.(*GroupMsgHandler)
		if b.groupMsgListeners == nil {
			b.groupMsgListeners = &groupMsgHeap{
				heap: []GroupMsgHandler{*a},
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
	}
	//其他信息
	if b.listeners[pos] == nil {
		b.listeners[pos] = &listenerHeap{
			heap: []listener{a},
		}
		return
	}
	b.listeners[pos].lock.Lock()
	defer b.listeners[pos].lock.Unlock()
	b.listeners[pos].Push(&a)
	sort.Sort(b.listeners[pos])
}
