package bot

import (
	"log"
)

// GroupRequest 邀请/申请加群都用这个
type GroupRequest struct {
	handle func(approve bool, flag, reason string) error
	flag   string
	// 邀请/申请加群的人的qq号
	UserID int64
	// 受邀/申请加入的群号
	GroupID int64
	// 验证消息
	Comment string
}

func (r *GroupRequest) Agree() {
	if err := r.handle(true, r.flag, ""); err != nil {
		log.Println(err)
	}
}

func (r *GroupRequest) Reject(reason string) {
	if err := r.handle(false, r.flag, reason); err != nil {
		log.Println(err)
	}
}

type FriendRequest struct {
	handle  func(approve bool, flag string, remark string) error
	flag    string
	UserID  int64
	Comment string
}

//Agree 同意加好友请求，remark为对好友的备注
func (r *FriendRequest) Agree(remark string) {
	if err := r.handle(true, r.flag, remark); err != nil {
		log.Println(err)
	}
}

func (r *FriendRequest) Reject() {
	if err := r.handle(false, r.flag, ""); err != nil {
		log.Println(err)
	}
}
