package api

import (
	"dingtalk/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// DingExtContact dingding extcontact
type DingExtContact struct {
	Tocken  string `json:"tocken" yaml:"tocken"`     // 应用访问tocken
	BaseURL string `json:"base_url" yaml:"base_url"` // 接口地址：https://oapi.dingtalk.com
}

// NewDingExtContact get DingExtContact object
func NewDingExtContact(baseURL string, tocken string) *DingExtContact {
	return &DingExtContact{
		BaseURL: baseURL,
		Tocken:  tocken,
	}
}

// ListLabelGroups 获取外部联系人标签列表
// 企业使用此接口可获取企业外部联系人标签列表，例如这个外部联系人是公司的客户，那么标签可能就是客户。
// size 分页大小，最大100
// offset 偏移位置
func (s *DingExtContact) ListLabelGroups(size int, offset int64) ([]*model.LabelGroups, error) {
	apiURL := s.BaseURL + "/topapi/extcontact/listlabelgroups?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		Size   int   `json:"size"`
		Offset int64 `json:"offset"`
	}{
		Size:   size,
		Offset: offset,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int                  `json:"errcode"`
		ErrMsg  string               `json:"errmsg"`
		Items   []*model.LabelGroups `json:"results"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// List 获取外部联系人列表
// size 分页大小，最大100
// offset 偏移位置
func (s *DingExtContact) List(size int, offset int64) ([]*model.ExtContact, error) {
	apiURL := s.BaseURL + "/topapi/extcontact/list?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		Size   int   `json:"size"`
		Offset int64 `json:"offset"`
	}{
		Size:   size,
		Offset: offset,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int                 `json:"errcode"`
		ErrMsg  string              `json:"errmsg"`
		Items   []*model.ExtContact `json:"results"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// Detail 获取外部联系人详情
// userID 用户id
func (s *DingExtContact) Detail(userID string) (*model.ExtContact, error) {
	apiURL := s.BaseURL + "/topapi/extcontact/list?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		UserID string `json:"user_id"`
	}{
		UserID: userID,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int               `json:"errcode"`
		ErrMsg  string            `json:"errmsg"`
		Items   *model.ExtContact `json:"result"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// Create 添加外部联系人
func (s *DingExtContact) Create(contact *model.ExtContactCreate) (string, error) {
	apiURL := s.BaseURL + "/topapi/extcontact/create?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		Contact *model.ExtContactCreate `json:"contact"`
	}{
		Contact: contact,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return "", err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		Items   string `json:"userid"` // 新外部联系人的userId
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return "", err
	}
	if result.ErrCode > 0 {
		return "", errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// Update 更新外部联系人
func (s *DingExtContact) Update(contact *model.ExtContactUpdate) error {
	apiURL := s.BaseURL + "/topapi/extcontact/update?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		Contact *model.ExtContactUpdate `json:"contact"`
	}{
		Contact: contact,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return err
	}
	if result.ErrCode > 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

// Delete 删除外部联系人
// userID 用户id
func (s *DingExtContact) Delete(userID string) error {
	apiURL := s.BaseURL + "/topapi/extcontact/delete?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		UserID string `json:"user_id"`
	}{
		UserID: userID,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return err
	}
	if result.ErrCode > 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}
