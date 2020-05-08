package api

import (
	"dingtalk/model"
	"fmt"
	"testing"
)

func getDingExtContact() *DingExtContact {
	tocken := "79616d2d98f83c72a280dc074a51102d"
	return NewDingExtContact(BASEURL, tocken)
}
func TestListLabelGroups(t *testing.T) {
	ext := getDingExtContact()
	size := 10
	var offset int64 = 0
	items, err := ext.ListLabelGroups(size, offset)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(items)
}

func TestExtContactList(t *testing.T) {
	ext := getDingExtContact()
	size := 10
	var offset int64 = 0
	items, err := ext.List(size, offset)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(items)
}

func TestExtContactDetail(t *testing.T) {
	ext := getDingExtContact()
	var userID string = "manager588"
	item, err := ext.Detail(userID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(item)
}

func TestExtContactCreate(t *testing.T) {
	ext := getDingExtContact()
	contact := model.ExtContactCreate{}
	item, err := ext.Create(&contact)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(item)
}

func TestExtContactUpdate(t *testing.T) {
	ext := getDingExtContact()
	contact := model.ExtContactUpdate{}
	if err := ext.Update(&contact); err != nil {
		t.Error(err)
		return
	}
}

func TestExtContactDelete(t *testing.T) {
	ext := getDingExtContact()
	userID := ""
	if err := ext.Delete(userID); err != nil {
		t.Error(err)
		return
	}
}
