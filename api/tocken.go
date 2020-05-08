package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

// DingTocken get dingding tocken parameters
type DingTocken struct {
	AppKey    string `json:"appKey" yaml:"appKey"`       // 应用的唯一标识key
	AppSecret string `json:"appSecret" yaml:"appSecret"` // 应用的密钥
	BaseURL   string `json:"base_url" yaml:"base_url"`   // 接口地址：https://oapi.dingtalk.com
}

// NewDingTocken get DingTocken object
func NewDingTocken(baseURL string, appKey string, appSecret string) *DingTocken {
	return &DingTocken{
		BaseURL:   baseURL,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

// GetTocken 获取tocken
// 正常情况下access_token有效期为7200秒，有效期内重复获取返回相同结果，并自动续期
func (s *DingTocken) GetTocken() (string, error) {
	apiURL := fmt.Sprintf("%s/gettoken?appkey=%s&appsecret=%s", s.BaseURL, s.AppKey, s.AppSecret)
	body, err := Get(apiURL)
	if err != nil {
		return "", err
	}
	result := struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return "", err
	}
	if result.ErrCode > 0 {
		return "", errors.New(result.ErrMsg)
	}
	return result.AccessToken, nil
}
