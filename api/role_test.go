package api

import (
	"fmt"
	"testing"
)

func getRole() *DingRole {
	tocken := "9dabccc23ddb37aab2964c3ef4ead528"
	return NewDingRole(BASEURL, tocken)
}

func TestRoleList(t *testing.T) {
	role := getRole()
	result, err := role.List(nil, nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result.HasMore)
}

func TestSimpleList(t *testing.T) {
	role := getRole()
	var roleID int64 = 0
	result, err := role.SimpleList(roleID, nil, nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result.HasMore)
}

func TestGetRoleGroup(t *testing.T) {
	role := getRole()
	var groupID int64 = 0
	result, err := role.GetRoleGroup(groupID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result.GroupName)
}

func TestGetRole(t *testing.T) {
	role := getRole()
	var roleID int64 = 0
	result, err := role.GetRole(roleID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result.Name)
}

func TestAddRole(t *testing.T) {
	role := getRole()
	var groupID int64 = 0
	var roleName string = ""
	result, err := role.AddRole(roleName, groupID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result)
}

func TestUpdateRole(t *testing.T) {
	role := getRole()
	var roleID int64 = 0
	var roleName string = ""
	if err := role.UpdateRole(roleName, roleID); err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteRole(t *testing.T) {
	role := getRole()
	var roleID int64 = 0
	if err := role.DeleteRole(roleID); err != nil {
		t.Error(err)
		return
	}
}

func TestAddRoleGroup(t *testing.T) {
	role := getRole()
	var name string = ""
	result, err := role.AddRoleGroup(name)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result)
}

func TestAddRolesForEmps(t *testing.T) {
	role := getRole()
	var roleIDs string = ""
	var userIDs string = ""
	if err := role.AddRolesForEmps(roleIDs, userIDs); err != nil {
		t.Error(err)
		return
	}
}

func TestRemoveRolesForEmps(t *testing.T) {
	role := getRole()
	var roleIDs string = ""
	var userIDs string = ""
	if err := role.RemoveRolesForEmps(roleIDs, userIDs); err != nil {
		t.Error(err)
		return
	}
}

func TestRoleScope(t *testing.T) {
	role := getRole()
	var roleID int64 = 0
	var userID string = ""
	if err := role.RoleScope(userID, roleID, nil); err != nil {
		t.Error(err)
		return
	}
}
