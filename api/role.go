package api

import (
	"dingtalk/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// DingRole dingding role
type DingRole struct {
	Tocken  string `json:"tocken" yaml:"tocken"`     // 应用访问tocken
	BaseURL string `json:"base_url" yaml:"base_url"` // 接口地址：https://oapi.dingtalk.com
}

// NewDingRole get DingRole object
func NewDingRole(baseURL string, tocken string) *DingRole {
	return &DingRole{
		BaseURL: baseURL,
		Tocken:  tocken,
	}
}

// List 获取角色列表
// size 分页大小，默认值：20，最大值200
// offset 分页偏移，默认值：0
func (s *DingRole) List(size *int, offset *int64) (*model.RoleResult, error) {
	apiURL := s.BaseURL + "/topapi/role/list?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	if size != nil {
		query.Set("size", strconv.Itoa(*size))
	}
	if offset != nil {
		query.Set("offset", strconv.FormatInt(*offset, 10))
	}
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int               `json:"errcode"`
		ErrMsg  string            `json:"errmsg"`
		Items   *model.RoleResult `json:"result"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// SimpleList 获取角色下的员工列表
// roleID 角色ID
// size 分页大小，默认值：20，最大值200
// offset 分页偏移，默认值：0
func (s *DingRole) SimpleList(roleID int64, size *int, offset *int64) (*model.RoleSimpleResult, error) {
	apiURL := s.BaseURL + "/topapi/role/simplelist?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("role_id", strconv.FormatInt(roleID, 10))
	if size != nil {
		query.Set("size", strconv.Itoa(*size))
	}
	if offset != nil {
		query.Set("offset", strconv.FormatInt(*offset, 10))
	}
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int                     `json:"errcode"`
		ErrMsg  string                  `json:"errmsg"`
		Items   *model.RoleSimpleResult `json:"result"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// GetRoleGroup 获取角色组
// groupID 角色组的Id
func (s *DingRole) GetRoleGroup(groupID int64) (*model.RoleGroup, error) {
	apiURL := s.BaseURL + "/topapi/role/simplelist?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("group_id", strconv.FormatInt(groupID, 10))
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int              `json:"errcode"`
		ErrMsg  string           `json:"errmsg"`
		Items   *model.RoleGroup `json:"role_group"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// GetRole 获取角色详情
// roleID 角色Id
func (s *DingRole) GetRole(roleID int64) (*model.RoleDetail, error) {
	apiURL := s.BaseURL + "/topapi/role/getrole?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("roleId", strconv.FormatInt(roleID, 10))
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int               `json:"errcode"`
		ErrMsg  string            `json:"errmsg"`
		Items   *model.RoleDetail `json:"role"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// AddRole 创建角色
// roleName 角色名称
// groupID 角色组id
func (s *DingRole) AddRole(roleName string, groupID int64) (int64, error) {
	apiURL := s.BaseURL + "/role/add_role?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		RoleName string `json:"roleName"`
		GroupID  int64  `json:"groupId"`
	}{
		RoleName: roleName,
		GroupID:  groupID,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return 0, err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		Items   int64  `json:"roleId"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return 0, err
	}
	if result.ErrCode > 0 {
		return 0, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// UpdateRole 更新角色
// roleName   角色名称
// roleID     角色id
func (s *DingRole) UpdateRole(roleName string, roleID int64) error {
	apiURL := s.BaseURL + "/role/update_role?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		RoleName string `json:"roleName"`
		RoleID   int64  `json:"roleId"`
	}{
		RoleName: roleName,
		RoleID:   roleID,
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

// DeleteRole 删除角色
func (s *DingRole) DeleteRole(roleID int64) error {
	apiURL := s.BaseURL + "/role/deleterole?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		RoleID int64 `json:"roleId"`
	}{
		RoleID: roleID,
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

// AddRoleGroup 创建角色组
// name 角色组名称
func (s *DingRole) AddRoleGroup(name string) (int64, error) {
	apiURL := s.BaseURL + "/role/add_role_group?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	body, err := Post(apiURL, bs)
	if err != nil {
		return 0, err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		GroupID int64  `json:"groupId"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return 0, err
	}
	if result.ErrCode > 0 {
		return 0, errors.New(result.ErrMsg)
	}
	return result.GroupID, nil
}

// AddRolesForEmps 批量增加员工角色
// roleIDs 角色id list，最大列表长度：20; 1,2,3,4,5
// userIDs 员工id list，最大列表长度：20; a,b,c,d,e
func (s *DingRole) AddRolesForEmps(roleIDs string, userIDs string) error {
	apiURL := s.BaseURL + "/role/add_role_group?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		RoleIds string `json:"roleIds"`
		UserIds string `json:"userIds"`
	}{
		RoleIds: roleIDs,
		UserIds: userIDs,
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

// RemoveRolesForEmps 批量删除员工角色
// roleIDs 角色id list，最大列表长度：20; 1,2
// userIDs 用户userId，最大列表长度：100; a,b
func (s *DingRole) RemoveRolesForEmps(roleIDs string, userIDs string) error {
	apiURL := s.BaseURL + "/role/removerolesforemps?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		RoleIds string `json:"roleIds"`
		UserIds string `json:"userIds"`
	}{
		RoleIds: roleIDs,
		UserIds: userIDs,
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

// RoleScope 设定角色成员管理范围
// userID 用户id
// roleID 角色id，必须是用户已经加入的角色
// deptIDs 部门id列表，最多50个，不传则设置范围为默认值：所有人
func (s *DingRole) RoleScope(userID string, roleID int64, deptIDs *string) error {
	apiURL := s.BaseURL + "/topapi/role/scope/update?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	data := struct {
		RoleID  int64   `json:"role_id"`
		UserID  string  `json:"userid"`
		DeptIds *string `json:"dept_ids"`
	}{
		RoleID:  roleID,
		UserID:  userID,
		DeptIds: deptIDs,
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
