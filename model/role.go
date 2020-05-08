package model

// RoleItem  角色
type RoleItem struct {
	ID        int64  `json:"id"`        // 角色id
	Name      string `json:"name"`      // 角色名称
	GroupName string `json:"groupName"` // 角色组名称
}

// RoleResult 角色列表查询接口返回结构
type RoleResult struct {
	HasMore bool `json:"hasMore"` // 是否还有更多数据
	List    []struct {
		Name    string `json:"name"`    // 角色组名称
		GroupID int64  `json:"groupId"` // 角色组id
		Roles   []struct {
			ID   int64  `json:"id"`   // 角色id
			Name string `json:"name"` // 角色名称
		} `json:"roles"`
	} `json:"list"` // 角色名称
}

// RoleSimpleResult 角色下的员工列表查询接口返回结构
type RoleSimpleResult struct {
	HasMore bool `json:"hasMore"` // 是否还有更多数据
	List    []struct {
		Name         string `json:"name"`   // 员工姓名
		UserID       string `json:"userid"` // 员工id
		ManageScopes []struct {
			DeptID int64  `json:"dept_id"` // 部门id
			Name   string `json:"name"`    // 部门名称
		} `json:"manageScopes"` // 管理范围
	} `json:"list"` // 角色名称
}

// RoleGroup 角色组查询接口返回结构
type RoleGroup struct {
	GroupName string `json:"group_name"` // 角色组名
	Roles     []struct {
		RoleID   int64  `json:"role_id"`   // 角色id
		RoleName string `json:"role_name"` // 角色名
	} `json:"roles"` // 角色列表信息
}

// RoleDetail  角色详情
type RoleDetail struct {
	Name    string `json:"name"`    // 角色名称
	GroupID int64  `json:"groupId"` // 角色组id
}
