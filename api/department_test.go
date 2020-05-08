package api

import (
	"dingtalk/model"
	"fmt"
	"testing"
)

func getDept() *DingDepartment {
	tocken := "9dabccc23ddb37aab2964c3ef4ead528"
	return NewDingDepartment(BASEURL, tocken)
}

func TestList(t *testing.T) {
	dept := getDept()
	fetchChild := true
	items, err := dept.List(nil, &fetchChild, nil)
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range items {
		s := fmt.Sprintf("id:%v,name:%s,parentid:%v,createDeptGroup:%v,autoAddUser:%v,ext:%s", item.ID, item.Name, item.ParentID, item.CreateDeptGroup, item.AutoAddUser, item.Ext)
		fmt.Println(s)
	}
}

func TestListIDs(t *testing.T) {
	dept := getDept()
	items, err := dept.ListIDs("1")
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range items {
		s := fmt.Sprintf("id:%v", item)
		fmt.Println(s)
	}
}

func TestDetail(t *testing.T) {
	dept := getDept()
	item, err := dept.Detail("344006982", nil)
	if err != nil {
		t.Error(err)
		return
	}
	s := fmt.Sprintf("id:%v,name:%s,parentid:%v,createDeptGroup:%v,autoAddUser:%v,ext:%s", item.ID, item.Name, item.ParentID, item.CreateDeptGroup, item.AutoAddUser, item.Ext)
	fmt.Println(s)
}

func TestListParentDeptsByDept(t *testing.T) {
	dept := getDept()
	item, err := dept.ListParentDeptsByDept("344006982")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(item)
}

func TestListParentDepts(t *testing.T) {
	dept := getDept()
	userID := ""
	items, err := dept.ListParentDepts(userID)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(items)
}

func TestDeptDelete(t *testing.T) {
	dept := getDept()
	deptID := "344439868"
	if err := dept.Delete(deptID); err != nil {
		t.Error(err)
		return
	}
}

func TestDeptUpdate(t *testing.T) {
	dept := getDept()
	name := "研发四组"
	d := model.DeptUdate{
		ID:   "344380429",
		Name: &name,
	}
	id, err := dept.Update(&d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(id)
}

func TestDeptCreate(t *testing.T) {
	dept := getDept()
	d := model.DeptCreate{
		Name:     "研发五组",
		ParentID: "344617093",
	}
	id, err := dept.Create(&d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(id)
}
