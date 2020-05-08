package api

import (
	"fmt"
	"testing"
)

func getUser() *DingUser {
	tocken := "79616d2d98f83c72a280dc074a51102d"
	return NewDingUser(BASEURL, tocken)
}

func TestUserDetail(t *testing.T) {
	user := getUser()
	userID := "manager588"
	ud, err := user.Detail(userID, nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ud)
}

func TestGetDeptMember(t *testing.T) {
	user := getUser()
	deptID := "1"
	ids, err := user.GetDeptMember(deptID)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(ids)
}

func TestSimplelist(t *testing.T) {
	user := getUser()
	var deptID int64 = 1
	uiList, err := user.Simplelist(deptID, nil, nil, nil, nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(uiList)
}

func TestListByPage(t *testing.T) {
	user := getUser()
	var deptID int64 = 1
	udList, err := user.ListByPage(deptID, nil, 0, 10, nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(udList)
}

func TestGetAdmin(t *testing.T) {
	user := getUser()
	adminList, err := user.GetAdmin()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(adminList)
}

func TestGetAdminScope(t *testing.T) {
	user := getUser()
	userID := "manager588"
	idList, err := user.GetAdminScope(userID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(idList)
}

func TestGetUseridByUnionid(t *testing.T) {
	user := getUser()
	unionid := "qmIDmr84iiXtHoJ4rQ33ORwiEiE"
	contactType, userID, err := user.GetUseridByUnionid(unionid)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("contactType:", contactType)
	fmt.Println("userID:", userID)
}

func TestGetByMobile(t *testing.T) {
	user := getUser()
	mobile := "18612522581"
	userID, err := user.GetByMobile(mobile)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("userID:", userID)
}

func TestGetOrgUserCount(t *testing.T) {
	user := getUser()
	onlyActive := 0
	count, err := user.GetOrgUserCount(onlyActive)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("count:", count)
}

func TestPostInactiveList(t *testing.T) {
	user := getUser()
	queryDate := "20200505"
	offset := 0
	size := 10
	hasMore, list, err := user.PostInactiveList(queryDate, offset, size)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("hasMore:", hasMore)
	fmt.Println("list:", list)
}

func TestUserDelete(t *testing.T) {
	user := getUser()
	userid := ""
	if err := user.Delete(userid); err != nil {
		t.Error(err)
		return
	}
	fmt.Println("删除成功")
}
