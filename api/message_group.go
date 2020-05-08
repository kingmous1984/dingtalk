package api

import (
	"dingtalk/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// 群会话消息是指可以调用接口创建企业群聊会话，然后可以以系统名义向群里推送群聊消息。

// DingMessageGroup dingding MessageGroup
type DingMessageGroup struct {
	Tocken  string `json:"tocken" yaml:"tocken"`     // 应用访问tocken
	BaseURL string `json:"base_url" yaml:"base_url"` // 接口地址：https://oapi.dingtalk.com
}

// NewDingMessageGroup get DingMessageGroup object
func NewDingMessageGroup(baseURL string, tocken string) *DingMessageGroup {
	return &DingMessageGroup{
		BaseURL: baseURL,
		Tocken:  tocken,
	}
}

// GetReadList 查询群消息已读人员列表
// messageID 发送群消息接口返回的加密消息id(消息id中包含url特殊字符时需要encode后再使用)
// cursor 分页查询的游标，第一次可以传0，后续传返回结果中的next_cursor的值。当返回结果中，没有next_cursor时，表示没有后续的数据了，可以结束调用
// size 分页查询的大小，最大可以传100
func (s *DingMessageGroup) GetReadList(messageID string, cursor int64, size int) (int64, []string, error) {
	apiURL := s.BaseURL + "/chat/getReadList?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("messageId", messageID)
	query.Set("cursor", strconv.FormatInt(cursor, 10))
	query.Set("size", strconv.Itoa(size))
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return 0, []string{}, err
	}
	result := struct {
		ErrCode        int      `json:"errcode"`
		ErrMsg         string   `json:"errmsg"`
		NextCursor     int64    `json:"next_cursor"`    // 下次分页获取的起始游标
		ReadUserIDList []string `json:"readUserIdList"` //已读人员的userId列表。已读人员为空时不返回该参数
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return 0, []string{}, err
	}
	if result.ErrCode > 0 {
		return 0, []string{}, errors.New(result.ErrMsg)
	}
	return result.NextCursor, result.ReadUserIDList, nil
}

// Create 创建会话
func (s *DingMessageGroup) Create(chat *model.ChatCreate) (string, int, error) {
	apiURL := s.BaseURL + "/chat/create?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	bs, err := json.Marshal(chat)
	if err != nil {
		return "", 0, err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return "", 0, err
	}
	result := struct {
		ErrCode         int    `json:"errcode"`
		ErrMsg          string `json:"errmsg"`
		ChatID          string `json:"chatid"`          // 群会话的id
		ConversationTag int    `json:"conversationTag"` // 会话类型。2表示企业群
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return "", 0, err
	}
	if result.ErrCode > 0 {
		return "", 0, errors.New(result.ErrMsg)
	}
	return result.ChatID, result.ConversationTag, nil
}

// Update 修改会话
func (s *DingMessageGroup) Update(chat *model.ChatUpdate) error {
	apiURL := s.BaseURL + "/chat/update?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	bs, err := json.Marshal(chat)
	if err != nil {
		return err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return err
	}
	result := struct {
		ErrCode         int    `json:"errcode"`
		ErrMsg          string `json:"errmsg"`
		ChatID          string `json:"chatid"`          // 群会话的id
		ConversationTag int    `json:"conversationTag"` // 会话类型。2表示企业群
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return err
	}
	if result.ErrCode > 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

// Get 获取会话
func (s *DingMessageGroup) Get(chatid string) (*model.ChatInfo, error) {
	apiURL := s.BaseURL + "/chat/get?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("chatid", chatid)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode  int             `json:"errcode"`
		ErrMsg   string          `json:"errmsg"`
		ChatInfo *model.ChatInfo `json:"chat_info"` // 群会话信息
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.ChatInfo, nil
}

// SendBase 基础发送
// cid 群会话的id，可以通过以下方式获取：
// （1）调用服务端API获取。调用创建群会话接口的返回chatid字段
// （2）调用前端API获取。小程序调用选择会话获取，H5微应用调用根据corpid选择会话获取。
func (s *DingMessageGroup) SendBase(cid string, bs []byte) (string, error) {
	apiURL := s.BaseURL + "/chat/send?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("chatid", cid)
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
// info 文本消息内容
func (s *DingMessageGroup) SendText(cid string, info string) (string, error) {
	bs, err := model.GetTextMessage(info)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}

// SendImage 发送图片消息
// mediaID 媒体文件id
func (s *DingMessageGroup) SendImage(cid string, mediaID string) (string, error) {
	bs, err := model.GetImageMessage(mediaID)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}

// SendVoice 发送音频消息
// mediaID 媒体文件id
func (s *DingMessageGroup) SendVoice(cid string, mediaID string, duration string) (string, error) {
	bs, err := model.GetVoiceMessage(mediaID, duration)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}

// SendFile 发送文件消息
// mediaID 媒体文件id
func (s *DingMessageGroup) SendFile(cid string, mediaID string) (string, error) {
	bs, err := model.GetFileMessage(mediaID)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}

// SendLink 发送链接消息
// link 链接结构参数
func (s *DingMessageGroup) SendLink(cid string, link *model.Link) (string, error) {
	bs, err := model.GetLinkMessage(link)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}

// SendOA 发送OA消息
// oa OA结构参数
func (s *DingMessageGroup) SendOA(cid string, oa *model.OA) (string, error) {
	bs, err := model.GetOAMessage(oa)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}

// SendMarkdown 发送Markdown消息
// markdown Markdown结构参数
func (s *DingMessageGroup) SendMarkdown(cid string, markdown *model.Markdown) (string, error) {
	bs, err := model.GetMarkdownMessage(markdown)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}

// SendActionCard 发送actionCard消息
// actionCard ActionCard结构参数
func (s *DingMessageGroup) SendActionCard(cid string, actionCard *model.ActionCard) (string, error) {
	bs, err := model.GetActionCardMessage(actionCard)
	if err != nil {
		return "", err
	}
	return s.SendBase(cid, bs)
}
