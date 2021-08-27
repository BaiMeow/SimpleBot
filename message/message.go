package message

import "strconv"

type ArrayMessage []arrayMessageUnit

type arrayMessageUnit struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type Msg []msgUnit

type msgUnit interface {
	//GetType 返回 text face image
	GetType() string
}

type Text struct {
	Text string
}

type Face struct {
	ID string
}
type Image struct {
	//发送时，file 参数可选已接受文件文件名，URL，base64编码
	File string
	Type string
	URL  string
}

type At struct {
	//ID 为qq号，ID为all时表示@全体成员
	ID string
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

func (u At) IsAt(id interface{}) bool {
	switch id.(type) {
	case string:
		return id.(string) == u.ID
	case int64:
		return strconv.FormatInt(id.(int64), 10) == u.ID

	}
	return false
}

func (a ArrayMessage) ToMsgStruct() Msg {
	var msg Msg
	for _, v := range a {
		switch v.Type {
		case "text":
			msg = append(msg, Text{
				Text: v.Data["text"].(string),
			})
		case "face":
			msg = append(msg, Face{
				ID: v.Data["id"].(string),
			})
		case "image":
			kind, ok := v.Data["type"].(string)
			if !ok {
				kind = ""
			}
			msg = append(msg, Image{
				File: v.Data["file"].(string),
				Type: kind,
				URL:  v.Data["url"].(string),
			})
		case "at":
			msg = append(msg, At{
				ID: v.Data["qq"].(string),
			})
		}
	}
	return msg
}

func (msg Msg) ToArrayMessage() ArrayMessage {
	var arrayMsg ArrayMessage
	for _, v := range msg {
		switch v.GetType() {
		case "text":
			tmp := v.(Text)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "text",
				Data: map[string]interface{}{
					"text": tmp.Text,
				},
			})
		case "face":
			tmp := v.(Face)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "face",
				Data: map[string]interface{}{
					"id": tmp.ID,
				},
			})
		case "image":
			tmp := v.(Image)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "image",
				Data: map[string]interface{}{
					"file": tmp.File,
					"type": tmp.Type,
					"url":  tmp.URL,
				},
			})
		case "at":
			tmp := v.(At)
			arrayMsg = append(arrayMsg, arrayMessageUnit{
				Type: "at",
				Data: map[string]interface{}{
					"qq": tmp.ID,
				},
			})
		}
	}
	return arrayMsg
}
