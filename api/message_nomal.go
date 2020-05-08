package api

import (
	"dingtalk/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// DingMessageNomal dingding MessageNomal
type DingMessageNomal struct {
	Tocken  string `json:"tocken" yaml:"tocken"`     // 应用访问tocken
	BaseURL string `json:"base_url" yaml:"base_url"` // 接口地址：https://oapi.dingtalk.com
}

// NewDingMessageNomal get DingMessageNomal object
func NewDingMessageNomal(baseURL string, tocken string) *DingMessageNomal {
	return &DingMessageNomal{
		BaseURL: baseURL,
		Tocken:  tocken,
	}
}

/**
普通会话消息是指员工个人在使用应用时，可以通过界面操作的方式往群或其他人的会话里推送消息，例如发送日志的场景。
发送普通消息，需要在前端页面调用JSAPI唤起联系人会话选择页面，选中后会返回会话cid，然后再调用服务端接口向会话里发送一条消息。

发送普通消息需要注意以下事项：
	不在当前接口调用所使用的企业的接收者（单聊接收者或者群聊会话里的接收者）不能收到消息。
	获取到的会话cid只能使用一次，且有效期为24小时。
*/

// SendBase 基础发送
func (s *DingMessageNomal) SendBase(sender string, cid string, bs []byte) (string, error) {
	apiURL := s.BaseURL + "/message/send_to_conversation?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("sender", sender)
	query.Set("cid", cid)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Post(apiURL, bs)
	if err != nil {
		return "", err
	}
	result := struct {
		ErrCode  int    `json:"errcode"`
		ErrMsg   string `json:"errmsg"`
		Receiver string `json:"receiver"` //接收者可以是单聊接收者或者群聊会话里的接收者，如果接收者是当前接口调用所使用的企业的员工，则是有效接收者。
		// 接口返回所有有效接收者的userId。非有效接收者是收不到消息的。"receiver": "UserID1|UserID2"
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return "", err
	}
	if result.ErrCode > 0 {
		return "", errors.New(result.ErrMsg)
	}
	return result.Receiver, nil
}

// SendText 发送文本消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// info 文本消息内容
func (s *DingMessageNomal) SendText(sender string, cid string, info string) (string, error) {
	bs, err := model.GetTextMessage(info)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}

// SendImage 发送图片消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// mediaID 媒体文件id
func (s *DingMessageNomal) SendImage(sender string, cid string, mediaID string) (string, error) {
	bs, err := model.GetImageMessage(mediaID)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}

// SendVoice 发送音频消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// mediaID 媒体文件id
func (s *DingMessageNomal) SendVoice(sender string, cid string, mediaID string, duration string) (string, error) {
	bs, err := model.GetVoiceMessage(mediaID, duration)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}

// SendFile 发送文件消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// mediaID 媒体文件id
func (s *DingMessageNomal) SendFile(sender string, cid string, mediaID string) (string, error) {
	bs, err := model.GetFileMessage(mediaID)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}

// SendLink 发送链接消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// link 链接结构参数
func (s *DingMessageNomal) SendLink(sender string, cid string, link *model.Link) (string, error) {
	bs, err := model.GetLinkMessage(link)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}

// SendOA 发送OA消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// oa OA结构参数
func (s *DingMessageNomal) SendOA(sender string, cid string, oa *model.OA) (string, error) {
	bs, err := model.GetOAMessage(oa)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}

// SendMarkdown 发送Markdown消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// markdown Markdown结构参数
func (s *DingMessageNomal) SendMarkdown(sender string, cid string, markdown *model.Markdown) (string, error) {
	bs, err := model.GetMarkdownMessage(markdown)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}

// SendActionCard 发送actionCard消息
// sender 消息发送者 userId
// cid 群会话或者个人会话的id，通过JSAPI接口唤起联系人界面选择会话获取会话cid；小程序参考获取会话信息，H5微应用参考获取会话信息
// actionCard ActionCard结构参数
func (s *DingMessageNomal) SendActionCard(sender string, cid string, actionCard *model.ActionCard) (string, error) {
	bs, err := model.GetActionCardMessage(actionCard)
	if err != nil {
		return "", err
	}
	return s.SendBase(sender, cid, bs)
}
