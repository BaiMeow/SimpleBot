package message

type ArrayMessage []arrayMessageUnit

type arrayMessageUnit struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type Msg []msgUnit

type msgUnit interface {
	getType() string
}

type Text struct {
	Text string
}

type Face struct {
	id string
}
type Image struct {
	//发送时，file 参数可选已接受文件文件名，URL，base64编码
	File string
	Type string
	URL  string
}

func (u Text) getType() string {
	return "text"
}

func (u Face) getType() string {
	return "face"
}

func (u Image) getType() string {
	return "image"
}

func (a *ArrayMessage) ToMsgStruct() *Msg {
	var msg Msg
	for _, v := range *a {
		switch v.Type {
		case "text":
			msg = append(msg, Text{
				Text: v.Data["text"].(string),
			})
		case "face":
			msg = append(msg, Face{
				id: v.Data["id"].(string),
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
		}
	}
	return &msg
}

func (a *Msg) ToArrayMessage() *ArrayMessage {
	var arrayMsg ArrayMessage
	for _, v := range *a {
		switch v.getType() {
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
					"id": tmp.id,
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
		}
	}
	return &arrayMsg
}
