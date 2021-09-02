package driver

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type wsDriver struct {
	addr  string
	token string
	conn  *websocket.Conn
}

func NewWsDriver(addr string, token string) *wsDriver {
	return &wsDriver{
		addr:  addr,
		token: token,
	}
}

func (d *wsDriver) Run() {
	header := http.Header{"Authorization": []string{"Bearer " + d.token}}
	conn, _, err := websocket.DefaultDialer.Dial(d.addr, header)
	if err != nil {
		log.Println("连接失败")
		time.Sleep(time.Second)
		d.Run()
		return
	}
	d.conn = conn
}

func (d *wsDriver) Write(data []byte) {
	for {
		if err := d.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Printf("无法将消息发送到onebot:%v", err)
			d.Run()
			continue
		}
		return
	}
}
func (d *wsDriver) Read() []byte {
	for {
		_, p, err := d.conn.ReadMessage()
		if err != nil {
			log.Printf("无法从onebot拉取消息:%v", err)
			d.Run()
			continue
		}
		return p
	}
}
func (d *wsDriver) Stop() {}
