package model

// LabelGroups 获取外部联系人标签列表接口返回结构
type LabelGroups struct {
	Name   string `json:"name"`  // 标签组名字
	Color  int64  `json:"color"` // 标签组颜色
	Labels []struct {
		ID   int64  `json:"id"`   // 标签id
		Name string `json:"name"` // 标签名字
	} `json:"labels"` // 标签列表
}

// ExtContact 外部联系人
type ExtContact struct {
	Title          string   `json:"title"`            // 职位
	ShareDeptIds   []int64  `json:"share_dept_ids"`   // 共享部门ID列表
	LabelIds       []int64  `json:"label_ids"`        // 标签
	Remark         string   `json:"title"`            // 备注
	Address        string   `json:"address"`          // 地址
	Name           string   `json:"name"`             // 姓名
	FollowerUserID string   `json:"follower_user_id"` // 负责人UserID
	StateCode      string   `json:"state_code"`       // 国家码
	CompanyName    string   `json:"company_name"`     // 公司名
	ShareUserIds   []string `json:"share_user_ids"`   // 共享员工UserID列表
	Mobile         string   `json:"mobile"`           // 手机号
	UserID         string   `json:"userid"`           // 外部联系人UserID
	Email          string   `json:"email"`            // 邮箱
}

// ExtContactCreate 外部联系人
type ExtContactCreate struct {
	Title          *string   `json:"title"`            // 职位
	ShareDeptIds   []*int64  `json:"share_dept_ids"`   // 共享部门ID列表
	LabelIds       []int64   `json:"label_ids"`        // 标签
	Remark         *string   `json:"title"`            // 备注
	Address        *string   `json:"address"`          // 地址
	Name           string    `json:"name"`             // 姓名
	FollowerUserID string    `json:"follower_user_id"` // 负责人UserID
	StateCode      string    `json:"state_code"`       // 国家码
	CompanyName    *string   `json:"company_name"`     // 公司名
	ShareUserIds   []*string `json:"share_user_ids"`   // 共享员工UserID列表
	Mobile         string    `json:"mobile"`           // 手机号
}

// ExtContactUpdate 外部联系人
type ExtContactUpdate struct {
	UserID         string    `json:"user_id"`          // 该外部联系人的userId
	Title          *string   `json:"title"`            // 职位
	ShareDeptIds   []*int64  `json:"share_dept_ids"`   // 共享部门ID列表
	LabelIds       []int64   `json:"label_ids"`        // 标签
	Remark         *string   `json:"title"`            // 备注
	Address        *string   `json:"address"`          // 地址
	Name           string    `json:"name"`             // 姓名
	FollowerUserID string    `json:"follower_user_id"` // 负责人UserID
	CompanyName    *string   `json:"company_name"`     // 公司名
	ShareUserIds   []*string `json:"share_user_ids"`   // 共享员工UserID列表
}
