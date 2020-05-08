package api

import (
	"dingtalk/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jinzhu/copier"
)

// DingDepartment dingding Department
type DingDepartment struct {
	Tocken  string `json:"tocken" yaml:"tocken"`     // 应用访问tocken
	BaseURL string `json:"base_url" yaml:"base_url"` // 接口地址：https://oapi.dingtalk.com
}

// NewDingDepartment get DingDepartment object
func NewDingDepartment(baseURL string, tocken string) *DingDepartment {
	return &DingDepartment{
		BaseURL: baseURL,
		Tocken:  tocken,
	}
}

// List 获取部门列表
// lang 通讯录语言（默认zh_CN，未来会支持en_US）
// fetchChild 是否递归部门的全部子部门
// 父部门id（如果不传，默认部门为根部门，根部门ID为1）
func (s *DingDepartment) List(lang *string, fetchChild *bool, id *string) ([]*model.Department, error) {
	apiURL := s.BaseURL + "/department/list?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	if lang != nil {
		query.Set("lang", *lang)
	}
	if fetchChild != nil {
		query.Set("fetch_child", strconv.FormatBool(*fetchChild))
	}
	if id != nil {
		query.Set("id", *id)
	}
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int                 `json:"errcode"`
		ErrMsg  string              `json:"errmsg"`
		Items   []*model.Department `json:"department"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// ListIDs 获取子部门ID列表
// id 父部门id。根部门的话传1
func (s *DingDepartment) ListIDs(id string) ([]int64, error) {
	apiURL := s.BaseURL + "/department/list_ids?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("id", id)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int     `json:"errcode"`
		ErrMsg  string  `json:"errmsg"`
		Items   []int64 `json:"sub_dept_id_list"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// Detail 获部门ID详情
// id 部门id
func (s *DingDepartment) Detail(id string, lang *string) (*model.DeptDetail, error) {
	var detail model.DeptDetail
	apiURL := s.BaseURL + "/department/get?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("id", id)
	if lang != nil {
		query.Set("lang", *lang)
	}
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		model.DeptDetail
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	copier.Copy(&detail, &result)
	return &detail, nil
}

// ListParentDeptsByDept 查询部门的所有上级父部门路径
// id 希望查询的部门的id，包含查询的部门本身
func (s *DingDepartment) ListParentDeptsByDept(id string) ([]int64, error) {
	apiURL := s.BaseURL + "/department/list_parent_depts_by_dept?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("id", id)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int     `json:"errcode"`
		ErrMsg  string  `json:"errmsg"`
		Items   []int64 `json:"parentIds"` // 该部门的所有父部门id列表
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// ListParentDepts 查询指定用户的所有上级父部门路径
// userId 希望查询的用户的id
func (s *DingDepartment) ListParentDepts(userId string) ([]int64, error) {
	apiURL := s.BaseURL + "/department/list_parent_depts?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("userId", userId)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int     `json:"errcode"`
		ErrMsg  string  `json:"errmsg"`
		Items   []int64 `json:"department"` // 该部门的所有父部门id列表
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// Delete 删除部门
// id 部门id(注：不能删除根部门；当部门里有员工，或者部门的子部门里有员工，也不能删除)
func (s *DingDepartment) Delete(id string) error {
	apiURL := s.BaseURL + "/department/delete?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("id", id)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
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

// Update 更新部门
func (s *DingDepartment) Update(dept *model.DeptUdate) (int64, error) {
	apiURL := s.BaseURL + "/department/update?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	bs, err := json.Marshal(dept)
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
		ID      int64  `json:"id"` // 已经更新的部门id
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return 0, err
	}
	if result.ErrCode > 0 {
		return 0, errors.New(result.ErrMsg)
	}
	return result.ID, nil
}

// Create 创建部门
func (s *DingDepartment) Create(dept *model.DeptCreate) (int64, error) {
	apiURL := s.BaseURL + "/department/create?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)

	apiURL += query.Encode()
	fmt.Println(apiURL)
	bs, err := json.Marshal(dept)
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
		ID      int64  `json:"id"` // 创建的部门id
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return 0, err
	}
	if result.ErrCode > 0 {
		return 0, errors.New(result.ErrMsg)
	}
	return result.ID, nil
}
