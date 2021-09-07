package message

import (
	"strconv"
	"strings"
)

type ArrayMessage []arrayMessageUnit

type arrayMessageUnit struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

// Msg 消息，是消息段的切片
type Msg []msgUnit

type msgUnit interface {
	// GetType 返回消息类型,
	// 现已支持 text image face at share group_contact user_contact
	GetType() string
}

type Text struct {
	Text string
}

type Face struct {
	ID string
}
type Image struct {
	// 发送时，file 参数可选已接受文件文件名，URL，base64编码
	File string
	Type string
	URL  string
}

type At struct {
	// ID 为qq号，ID为all时表示@全体成员
	ID string
}

//Share 链接分享
type Share struct {
	URL   string
	Title string
	//发送时可选，内容描述
	Content string
	//发送时可选，图片URL，部分onebot实现不一定有效
	Image string
}

//UserContact 推荐好友
type UserContact struct {
	ID int64
}

//GroupContact 推荐群
type GroupContact struct {
	ID int64
}

//Location 位置
type Location struct {
	Lat     float64
	Lon     float64
	Title   string
	Content string
}

type Reply struct {
	ID int32
}

func (u Text) GetType() string {
	return "text"
}

func (u Face) GetType() string {
	return "face"
}

func (u Image) GetType() string {
	return "image"
}

func (u At) GetType() string {
	return "at"
}

func (u Share) GetType() string {
	return "share"
}

func (u UserContact) GetType() string {
	return "user_contact"
}

func (u GroupContact) GetType() string {
	return "group_contact"
}

func (u Location) GetType() string {
	return "location"
}

func (u Reply) GetType() string {
	return "reply"
}

// IsAt 判断时候在@某人，支持int64和string，可以传入"all"表示@全体
func (u At) IsAt(id interface{}) bool {
	switch id.(type) {
	case string:
		return id.(string) == u.ID
	case int64:
		return strconv.FormatInt(id.(int64), 10) == u.ID
	}
	return false
}

// ToMsgStruct 将map数组消息转为结构体数组格式，一般不要使用
func (a ArrayMessage) ToMsgStruct() Msg {
	var msg Msg
	for _, v := range a {
		switch v.Type {
		case "text":
			msg = append(msg, Text{Text: v.Data["text"]})
		case "face":
			msg = append(msg, Face{ID: v.Data["id"]})
		case "image":
			msg = append(msg, Image{
				File: v.Data["file"],
				Type: v.Data["type"],
				URL:  v.Data["url"],
			})
		case "at":
			msg = append(msg, At{ID: v.Data["qq"]})
		case "share":
			msg = append(msg, Share{
				URL:     v.Data["url"],
				Title:   v.Data["title"],
				Content: v.Data["content"],
				Image:   v.Data["image"],
			})
		case "contact":
			id, err := strconv.ParseInt(v.Data["id"], 10, 64)
			if err != nil {
				continue
			}
			switch v.Data["type"] {
			case "qq":
				msg = append(msg, UserContact{ID: id})
			case "group":
				msg = append(msg, GroupContact{ID: id})
			}
		case "location":
			lat, err := strconv.ParseFloat(v.Data["lat"], 64)
			if err != nil {
				continue
			}
			lon, err := strconv.ParseFloat(v.Data["lon"], 64)
			if err != nil {
				continue
			}
			msg = append(msg, Location{
				Lat:     lat,
				Lon:     lon,
				Title:   v.Data["title"],
				Content: v.Data["content"],
			})
		case "reply":
			id, err := strconv.ParseInt(v.Data["id"], 10, 32)
			if err != nil {
				continue
			}
			msg = append(msg, Reply{ID: int32(id)})
		}
	}
	return msg
}

// ToArrayMessage 将结构体数组消息转为map数组，一般不要使用
func (msg Msg) ToArrayMessage() ArrayMessage {
	var arrayMsg ArrayMessage
	for _, v := range msg {
		switch v.GetType() {
		case "text":
			tmp := v.(Text)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "text",
				Data: map[string]string{
					"text": tmp.Text,
				},
			})
		case "face":
			tmp := v.(Face)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "face",
				Data: map[string]string{
					"id": tmp.ID,
				},
			})
		case "image":
			tmp := v.(Image)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "image",
				Data: map[string]string{
					"file": tmp.File,
					"type": tmp.Type,
					"url":  tmp.URL,
				},
			})
		case "at":
			tmp := v.(At)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "at",
				Data: map[string]string{
					"qq": tmp.ID,
				},
			})
		case "share":
			tmp := v.(Share)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "share",
				Data: map[string]string{
					"url":     tmp.URL,
					"title":   tmp.Title,
					"content": tmp.Content,
					"image":   tmp.Image,
				},
			})
		case "user_contact":
			tmp := v.(UserContact)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "contact",
				Data: map[string]string{
					"type": "qq",
					"id":   strconv.FormatInt(tmp.ID, 10),
				},
			})
		case "group_contact":
			tmp := v.(GroupContact)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "contact",
				Data: map[string]string{
					"type": "group",
					"id":   strconv.FormatInt(tmp.ID, 10),
				},
			})

		case "location":
			tmp := v.(Location)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "location",
				Data: map[string]string{
					"lat":     strconv.FormatFloat(tmp.Lat, 'f', -1, 64),
					"lon":     strconv.FormatFloat(tmp.Lon, 'f', -1, 64),
					"title":   tmp.Title,
					"content": tmp.Content,
				},
			})
		case "reply":
			tmp := v.(Reply)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "reply",
				Data: map[string]string{
					"id": strconv.FormatInt(int64(tmp.ID), 10),
				},
			})
		}

	}
	return arrayMsg
}

//Fields 将text消息以空格分割符划分成多个，顺序不变，并且保留其他特殊消息类型，建议用于命令解析的预处理
func (msg Msg) Fields() Msg {
	var newMsg Msg
	for _, unit := range msg {
		if unit.GetType() != "text" {
			newMsg = append(newMsg, unit)
			continue
		}
		txt := unit.(Text)
		args := strings.Fields(txt.Text)
		for _, arg := range args {
			newMsg = append(newMsg, Text{Text: arg})
		}
	}
	return newMsg
}
