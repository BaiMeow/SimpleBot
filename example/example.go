//main.go暂时作为示例和调试使用
package main

import (
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/message"
	"log"
)

const addr = "ws://localhost:6700"
const token = ""

var b *bot.Bot

func main() {
	//目前仅实现了正向WS，这里使用正向ws作为底层驱动
	b = bot.New(driver.NewWsDriver(addr, token))
	//添加事件处理er，注意本bot不支持匿名聊天
	//群消息
	b.Attach(&bot.GroupMsgHandler{
		Priority: 1,
		F:        justReply,
	})
	//私聊消息
	b.Attach(&bot.PrivateMsgHandler{
		Priority: 1,
		F:        justReply2,
	})
	//群成员添加
	b.Attach(&bot.GroupAddHandler{
		Priority: 1,
		F:        agree,
	})
	//群成员减少
	b.Attach(&bot.GroupDecreaseHandler{
		Priority: 1,
		F:        groupdecrease,
	})
	//执行后bot开始运行，会在控制台输出登陆的qq号
	b.Run()
	//这里bot已经在运行了，没什么要做的事情话就堵塞了吧
	select {}
}

//复读
func justReply(MsgID int32, GroupID int64, UserID int64, Msg message.Msg) bool {
	log.Println("new message")
	//将收到的消息发回去，实现复读
	if msgid, err := b.SendGroupMsg(GroupID, Msg); err != nil {
		log.Println(err)
	} else {
		log.Println(msgid)
	}
	return false
}

//对纯文本消息，回复指定内容
func justReply2(MsgID int32, UserID int64, msg message.Msg) bool {
	//不是纯文本就忽略，纯文本肯定只有一段文字而且这一段的类型为纯文本
	if len(msg) != 1 || msg[0].GetType() != "text" {
		return false
	}
	log.Println("new message")
	//发送指定内容
	if msgid, err := b.SendPrivateMsg(UserID,
		//构造一则消息，现在构造的这则消息由纯文本，图片，表情组成
		message.New().
			Text("BaiMeow").
			Image("https://baimeow.cn/me/BaiMeow.jpg", false).
			Face("1"),
	); err != nil {
		log.Println(err)
	} else {
		log.Println(msgid)
	}
	return false
}

//同意所有入群申请，为了方便同意所以这里的参数不是散装的
func agree(request *bot.GroupRequest) bool {
	log.Println(request)
	//同意
	request.Agree()
	return true
}

//群人数减少
func groupdecrease(GroupID, OperatorID, UserID int64) bool {
	log.Println("-1s:", GroupID, OperatorID, UserID)
	return true
}
