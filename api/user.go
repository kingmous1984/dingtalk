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

// DingUser dingding user
type DingUser struct {
	Tocken  string `json:"tocken" yaml:"tocken"`     // 应用访问tocken
	BaseURL string `json:"base_url" yaml:"base_url"` // 接口地址：https://oapi.dingtalk.com
}

// NewDingUser get DingUser object
func NewDingUser(baseURL string, tocken string) *DingUser {
	return &DingUser{
		BaseURL: baseURL,
		Tocken:  tocken,
	}
}

// Detail 获取用户详情
// id 部门id
func (s *DingUser) Detail(userid string, lang *string) (*model.UserRoles, error) {
	var detail model.UserRoles
	apiURL := s.BaseURL + "/user/get?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("userid", userid)
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
		model.UserDetail
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

// GetDeptMember 获取部门用户userid列表
// deptId 部门id
func (s *DingUser) GetDeptMember(deptID string) ([]string, error) {
	apiURL := s.BaseURL + "/user/getDeptMember?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("deptId", deptID)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int      `json:"errcode"`
		ErrMsg  string   `json:"errmsg"`
		Items   []string `json:"userIds"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// Simplelist 获取部门用户
// deptID 获取的部门id
// lang 通讯录语言(默认zh_CN另外支持en_US)
// offset 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量
// size 支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
// order 支持分页查询，部门成员的排序规则，默认不传是按自定义排序；
// entry_asc：代表按照进入部门的时间升序，
// entry_desc：代表按照进入部门的时间降序，
// modify_asc：代表按照部门信息修改时间升序，
// modify_desc：代表按照部门信息修改时间降序，
// custom：代表用户定义(未定义时按照拼音)排序
func (s *DingUser) Simplelist(deptID int64, lang *string, offset *int64, size *int64, order *string) (*model.UserItemsList, error) {
	var list model.UserItemsList
	apiURL := s.BaseURL + "/user/simplelist?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("department_id", strconv.FormatInt(deptID, 10))
	if lang != nil {
		query.Set("lang", *lang)
	}
	if lang != nil {
		query.Set("offset", strconv.FormatInt(*offset, 10))
	}
	if offset != nil {
		query.Set("size", strconv.FormatInt(*offset, 10))
	}
	if size != nil {
		query.Set("size", strconv.FormatInt(*size, 10))
	}
	if order != nil {
		query.Set("order", *order)
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
		model.UserItemsList
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	copier.Copy(&list, &result)
	return &list, nil
}

// ListByPage 获取部门用户详情
// deptID 获取的部门id
// lang 通讯录语言(默认zh_CN另外支持en_US)
// offset 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量
// size 支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
// order 支持分页查询，部门成员的排序规则，默认不传是按自定义排序；
// entry_asc：代表按照进入部门的时间升序，
// entry_desc：代表按照进入部门的时间降序，
// modify_asc：代表按照部门信息修改时间升序，
// modify_desc：代表按照部门信息修改时间降序，
// custom：代表用户定义(未定义时按照拼音)排序
func (s *DingUser) ListByPage(deptID int64, lang *string, offset int64, size int64, order *string) (*model.UserDetailList, error) {
	var list model.UserDetailList
	apiURL := s.BaseURL + "/user/listbypage?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("department_id", strconv.FormatInt(deptID, 10))
	query.Set("offset", strconv.FormatInt(offset, 10))
	query.Set("size", strconv.FormatInt(offset, 10))
	query.Set("size", strconv.FormatInt(size, 10))
	if lang != nil {
		query.Set("lang", *lang)
	}
	if order != nil {
		query.Set("order", *order)
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
		model.UserDetailList
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	copier.Copy(&list, &result)
	return &list, nil
}

// GetAdmin 获取管理员列表
func (s *DingUser) GetAdmin() ([]*model.AdminItem, error) {
	apiURL := s.BaseURL + "/user/get_admin?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int                `json:"errcode"`
		ErrMsg  string             `json:"errmsg"`
		Items   []*model.AdminItem `json:"admin_list"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// GetAdminScope 获取管理员通讯录权限范围
// userid 员工id
func (s *DingUser) GetAdminScope(userid string) ([]int64, error) {
	apiURL := s.BaseURL + "/topapi/user/get_admin_scope?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("userid", userid)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		ErrCode int     `json:"errcode"`
		ErrMsg  string  `json:"errmsg"`
		Items   []int64 `json:"dept_ids"` // 可管理的部门id列表
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return nil, err
	}
	if result.ErrCode > 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.Items, nil
}

// GetUseridByUnionid 根据unionid获取userid
// unionid 员工在当前开发者企业账号范围内的唯一标识，系统生成，固定值，不会改变
func (s *DingUser) GetUseridByUnionid(unionid string) (int, string, error) {
	apiURL := s.BaseURL + "/user/getUseridByUnionid?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("unionid", unionid)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return 0, "", err
	}
	result := struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		ContactType int    `json:"contactType"` // 联系类型，0表示企业内部员工，1表示企业外部联系人
		UserID      string `json:"userid"`      // 员工id
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return 0, "", err
	}
	if result.ErrCode > 0 {
		return 0, "", errors.New(result.ErrMsg)
	}
	return result.ContactType, result.UserID, nil
}

// GetByMobile 根据手机号获取userid
// mobile 手机号码
func (s *DingUser) GetByMobile(mobile string) (string, error) {
	apiURL := s.BaseURL + "/user/get_by_mobile?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("mobile", mobile)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return "", err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		UserID  string `json:"userid"` // 员工在当前企业内的唯一标识。
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return "", err
	}
	if result.ErrCode > 0 {
		return "", errors.New(result.ErrMsg)
	}
	return result.UserID, nil
}

// GetOrgUserCount 获取企业员工人数
// onlyActive 0：包含未激活钉钉的人员数量; 1：不包含未激活钉钉的人员数量
func (s *DingUser) GetOrgUserCount(onlyActive int) (int, error) {
	apiURL := s.BaseURL + "/user/get_org_user_count?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("onlyActive", strconv.Itoa(onlyActive))
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Get(apiURL)
	if err != nil {
		return 0, err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		Count   int    `json:"count"` // 企业员工数量
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return 0, err
	}
	if result.ErrCode > 0 {
		return 0, errors.New(result.ErrMsg)
	}
	return result.Count, nil
}

// PostInactiveList 未登录钉钉的员工列表
// queryDate 查询日期: 20190808
// offset 分页数据偏移量，从0开始
// size 每页大小，最大100
func (s *DingUser) PostInactiveList(queryDate string, offset int, size int) (bool, []string, error) {
	apiURL := s.BaseURL + "/topapi/inactive/user/get?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("query_date", queryDate)
	query.Set("offset", strconv.Itoa(offset))
	query.Set("size", strconv.Itoa(size))
	apiURL += query.Encode()
	fmt.Println(apiURL)
	body, err := Post(apiURL, []byte{})
	if err != nil {
		return false, []string{}, err
	}
	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		Result  struct {
			HasMore bool     `json:"has_more"` // 是否还有更多数据
			List    []string `json:"list"`     // 未登录用户userId列表
		} `json:"result"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return false, []string{}, err
	}
	if result.ErrCode > 0 {
		return false, []string{}, errors.New(result.ErrMsg)
	}
	return result.Result.HasMore, result.Result.List, nil
}

// Delete 删除用户
// userid 员工id
func (s *DingUser) Delete(userid string) error {
	apiURL := s.BaseURL + "/user/delete?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	query.Set("userid", userid)
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

// Update 更新用户
func (s *DingUser) Update(user *model.User) error {
	apiURL := s.BaseURL + "/user/update?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	bs, err := json.Marshal(user)
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
		Result  struct {
			HasMore bool     `json:"has_more"` // 是否还有更多数据
			List    []string `json:"list"`     // 未登录用户userId列表
		} `json:"result"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return err
	}
	if result.ErrCode > 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

// Create 创建用户
func (s *DingUser) Create(user *model.User) error {
	apiURL := s.BaseURL + "/user/create?"
	query := url.Values{}
	query.Set("access_token", s.Tocken)
	apiURL += query.Encode()
	fmt.Println(apiURL)
	bs, err := json.Marshal(user)
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
		Result  struct {
			HasMore bool     `json:"has_more"` // 是否还有更多数据
			List    []string `json:"list"`     // 未登录用户userId列表
		} `json:"result"`
	}{}
	if err := json.Unmarshal(body.([]byte), &result); err != nil {
		return err
	}
	if result.ErrCode > 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}
