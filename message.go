// Package cross build Lark messages and output to chyroc/lark format
package cross

import (
	"encoding/json"

	clark "github.com/chyroc/lark"
	golark "github.com/go-lark/lark/v2"
)

// BuildMessage .
func BuildMessage(receiveIDType string, om golark.OutcomingMessage) (*clark.SendRawMessageReq, error) {
	req := &clark.SendRawMessageReq{
		ReceiveIDType: clark.IDType(receiveIDType),
		ReceiveID:     buildReceiveID(om),
		MsgType:       clark.MsgType(om.MsgType),
		Content:       buildContent(om),
	}
	if req.ReceiveID == "" {
		return nil, golark.ErrInvalidReceiveID
	}
	if req.Content == "" {
		return nil, golark.ErrMessageNotBuild
	}

	if om.UUID != "" {
		req.UUID = &om.UUID
	}

	return req, nil
}

func buildContent(om golark.OutcomingMessage) string {
	var (
		content = ""
		b       []byte
		err     error
	)
	switch om.MsgType {
	case golark.MsgText:
		b, err = json.Marshal(om.Content.Text)
	case golark.MsgImage:
		b, err = json.Marshal(om.Content.Image)
	case golark.MsgFile:
		b, err = json.Marshal(om.Content.File)
	case golark.MsgShareCard:
		b, err = json.Marshal(om.Content.ShareChat)
	case golark.MsgShareUser:
		b, err = json.Marshal(om.Content.ShareUser)
	case golark.MsgPost:
		b, err = json.Marshal(om.Content.Post)
	case golark.MsgInteractive:
		if om.Content.Card != nil {
			b, err = json.Marshal(om.Content.Card)
		} else if om.Content.Template != nil {
			b, err = json.Marshal(om.Content.Template)
		}
	case golark.MsgAudio:
		b, err = json.Marshal(om.Content.Audio)
	case golark.MsgMedia:
		b, err = json.Marshal(om.Content.Media)
	case golark.MsgSticker:
		b, err = json.Marshal(om.Content.Sticker)
	}
	if err != nil {
		return ""
	}
	content = string(b)

	return content
}

func buildReceiveID(om golark.OutcomingMessage) string {
	switch om.UIDType {
	case golark.UIDEmail:
		return om.Email
	case golark.UIDUserID:
		return om.UserID
	case golark.UIDOpenID:
		return om.OpenID
	case golark.UIDChatID:
		return om.ChatID
	case golark.UIDUnionID:
		return om.UnionID
	}
	return ""
}
