package api

import (
	"fmt"
	"testing"
)

const BASEURL = "https://oapi.dingtalk.com"

func TestGetTocken(t *testing.T) {
	appKey := ""
	appSecret := ""
	dingTocken := NewDingTocken(BASEURL, appKey, appSecret)
	tocken, err := dingTocken.GetTocken()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tocken)
}
