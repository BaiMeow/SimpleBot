package bot

import "log"

type GroupAddRequest struct {
	handler func(approve bool, flag, reason string) error
	flag    string
	//申请加群的人的qq号
	UserID int64
	//申请加入的群号
	GroupID int64
	//验证消息
	Comment string
}

func (r *GroupAddRequest) Agree() {
	if err := r.handler(true, r.flag, ""); err != nil {
		log.Println(err)
	}
}

func (r *GroupAddRequest) Reject(reason string) {
	if err := r.handler(false, r.flag, reason); err != nil {
		log.Println(err)
	}
}
