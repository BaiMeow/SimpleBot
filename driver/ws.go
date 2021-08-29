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
	if err := d.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Fatal(err)
	}
}
func (d *wsDriver) Read() []byte {
	_, p, err := d.conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	return p
}
func (d *wsDriver) Stop() {}
