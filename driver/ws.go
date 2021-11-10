package driver

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var ErrConnClosed = errors.New("connection was closed")

type wsDriver struct {
	addr   string
	token  string
	conn   *websocket.Conn
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
		if err := d.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Printf("无法将消息发送到onebot:%v", err)
			if d.closed {
				return ErrConnClosed
			}
			if err := d.Run(); err != nil {
				return err
			}
			continue
		}
		return nil
	}
}
func (d *wsDriver) Read() ([]byte, error) {
	for {
		_, p, err := d.conn.ReadMessage()
		if err != nil {
			log.Printf("无法从onebot拉取消息:%v", err)
			if d.closed {
				return nil, nil
			}
			if err := d.Run(); err != nil {
				return nil, err
			}
			continue
		}
		return p, nil
	}
}
func (d *wsDriver) Stop() error {
	d.closed = true
	return d.conn.Close()
}
