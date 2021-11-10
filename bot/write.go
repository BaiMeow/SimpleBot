package bot

import (
	"errors"
	"github.com/BaiMeow/SimpleBot/driver"
	"log"
	"time"
)

func (b *Bot) writeWithRetry(data []byte) error {
	for {
		err := b.driver.Write(data)
		if err == nil {
			return nil
		}
		//重试三次
		for i := 0; i < 3 && err != nil; i++ {
			log.Println(err)
			if errors.Is(err, driver.ErrConnClosed) {
				return err
			}
			time.Sleep(time.Second)
			err = b.driver.Run()
		}
		if err != nil {
			return err
		}
	}
}
