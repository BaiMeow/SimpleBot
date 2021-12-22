package driver

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type wsDriver struct {
	addr   string
	token  string
	conn   *websocket.Conn
	muR    sync.Mutex
	muW    sync.Mutex
	closed bool
}

func NewWsDriver(addr string, token string) *wsDriver {
	return &wsDriver{
		addr:   addr,
		token:  token,
		closed: true,
	}
}

func (d *wsDriver) Run() error {
	if !d.closed {
		if err := d.Stop(); err != nil {
			return err
		}
	}
	header := http.Header{"Authorization": []string{"Bearer " + d.token}}
	var (
		err  error
		conn *websocket.Conn
	)
	//尝试三次
	for i := 0; i < 3; i++ {
		conn, _, err = websocket.DefaultDialer.Dial(d.addr, header)
		if err == nil {
			d.conn = conn
			d.closed = false
			return nil
		}
	}
	return err
}

func (d *wsDriver) Write(data []byte) error {
	for {
		d.muW.Lock()
		err := d.conn.WriteMessage(websocket.TextMessage, data)
		d.muW.Unlock()
		if err == nil {
			return nil
		}
		log.Printf("无法将消息发送到onebot:%v", err)
		if d.closed {
			return ErrConnClosed
		}
		d.Stop()
		if err := d.Run(); err != nil {
			return err
		}
	}
}
func (d *wsDriver) Read() ([]byte, error) {
	for {
		d.muR.Lock()
		_, p, err := d.conn.ReadMessage()
		d.muR.Unlock()
		if err == nil {
			return p, nil
		}
		log.Printf("无法从onebot拉取消息:%v", err)
		if d.closed {
			return nil, nil
		}
		d.Stop()
		if err := d.Run(); err != nil {
			return nil, err
		}
	}
}
func (d *wsDriver) Stop() error {
	d.closed = true
	return d.conn.Close()
}
