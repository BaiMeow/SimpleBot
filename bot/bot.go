package bot

import (
	"sort"
	"sync"

	"github.com/BaiMeow/SimpleBot/driver"
)

type Bot struct {
	driver driver.Driver

	listeners map[string]*listenerHeap
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

func (b *Bot) Attach(pos string, a listener) {
	//todo:check listener valid
	if b.listeners[pos] == nil {
		b.listeners[pos] = &listenerHeap{
			heap: []listener{a},
		}
		return
	}
	b.listeners[pos].lock.Lock()
	defer b.listeners[pos].lock.Unlock()
	b.listeners[pos].Push(a)
	sort.Sort(b.listeners[pos])
}

type listenerHeap struct {
	heap []listener
	lock sync.Mutex
}
type listener interface {
	GetPriority() int
}

func (h *listenerHeap) Len() int { return len(h.heap) }
func (h *listenerHeap) Less(i, j int) bool {
	return h.heap[i].GetPriority() < h.heap[j].GetPriority()
}
func (h *listenerHeap) Swap(i, j int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}
func (h *listenerHeap) Push(l listener) {
	h.heap = append(h.heap, l)
}
