package model

// Department 部门列表
type Department struct {
	ID              int64  `json:"id"`              // 部门id
	Name            string `json:"name"`            // 部门名称
	ParentID        int64  `json:"parentid"`        // 父部门id，根部门为1
	CreateDeptGroup bool   `json:"createDeptGroup"` // 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	AutoAddUser     bool   `json:"autoAddUser"`     // 当群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	Ext             JSON   `json:"ext"`             // 部门自定义字段
}

// DeptDetail  部门详情
type DeptDetail struct {
	ID                    int64  `json:"id"`                    // 部门id
	Name                  string `json:"name"`                  // 部门名称
	ParentID              int64  `json:"parentid"`              // 父部门id，根部门为1
	Order                 int    `json:"order"`                 // 当前部门在父部门下的所有子部门中的排序值
	CreateDeptGroup       bool   `json:"createDeptGroup"`       // 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	AutoAddUser           bool   `json:"autoAddUser"`           // 当群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	DeptHiding            bool   `json:"deptHiding"`            // 是否隐藏部门，true表示隐藏，false表示显示
	DeptPermits           string `json:"deptPermits"`           // 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的的字符串，使用“|”符号进行分割
	UserPermits           string `json:"userPermits"`           // 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的的字符串，使用“|”符号进行分割
	OuterDept             bool   `json:"outerDept"`             // 是否本部门的员工仅可见员工自己，为true时，本部门员工默认只能看到员工自己
	OuterPermitDepts      string `json:"outerPermitDepts"`      // 本部门的员工仅可见员工自己为true时，可以配置额外可见部门，值为部门id组成的的字符串，使用“|”符号进行分割
	OuterPermitUsers      string `json:"outerPermitUsers"`      // 本部门的员工仅可见员工自己为true时，可以配置额外可见人员，值为userid组成的的字符串，使用“|”符号进行分割
	OrgDeptOwner          string `json:"orgDeptOwner"`          // 企业群群主
	DeptManagerUseridList string `json:"deptManagerUseridList"` // 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用“|”符号进行分割
	SourceIdentifier      string `json:"sourceIdentifier"`      // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	GroupContainSubDept   bool   `json:"groupContainSubDept"`   // 部门群是否包含子部门
	Ext                   JSON   `json:"ext"`                   // 部门自定义字段
}

// DeptUdate  部门编辑
type DeptUdate struct {
	ID                     string  `json:"id"`                     // 部门id
	Lang                   *string `json:"lang"`                   // 通讯录语言(默认zh_CN另外支持en_US)
	Name                   *string `json:"name"`                   // 部门名称
	ParentID               *string `json:"parentid"`               // 父部门id，根部门为1
	Order                  *string `json:"order"`                  // 当前部门在父部门下的所有子部门中的排序值
	CreateDeptGroup        *bool   `json:"createDeptGroup"`        // 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	AutoAddUser            *bool   `json:"autoAddUser"`            // 当群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	DeptHiding             *bool   `json:"deptHiding"`             // 是否隐藏部门，true表示隐藏，false表示显示
	DeptPermits            *string `json:"deptPermits"`            // 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的的字符串，使用“|”符号进行分割
	UserPermits            *string `json:"userPermits"`            // 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的的字符串，使用“|”符号进行分割
	OuterDept              *bool   `json:"outerDept"`              // 是否本部门的员工仅可见员工自己，为true时，本部门员工默认只能看到员工自己
	OuterPermitDepts       *string `json:"outerPermitDepts"`       // 本部门的员工仅可见员工自己为true时，可以配置额外可见部门，值为部门id组成的的字符串，使用“|”符号进行分割
	OuterPermitUsers       *string `json:"outerPermitUsers"`       // 本部门的员工仅可见员工自己为true时，可以配置额外可见人员，值为userid组成的的字符串，使用“|”符号进行分割
	OuterDeptOnlySelf      *bool   `json:"outerDeptOnlySelf"`      // outerDept为true时，可以配置该字段，为true时，表示只能看到所在部门及下级部门通讯录
	OrgDeptOwner           *string `json:"orgDeptOwner"`           // 企业群群主
	DeptManagerUseridList  *string `json:"deptManagerUseridList"`  // 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用“|”符号进行分割
	SourceIdentifier       *string `json:"sourceIdentifier"`       // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	GroupContainSubDept    *bool   `json:"groupContainSubDept"`    // 部门群是否包含子部门
	GroupContainOuterDept  *bool   `json:"groupContainOuterDept"`  // 部门群是否包含外包部门
	GroupContainHiddenDept *bool   `json:"groupContainHiddenDept"` // 部门群是否包含隐藏部门
	Ext                    *JSON   `json:"ext"`                    // 部门自定义字段
}

// DeptCreate  部门新建
type DeptCreate struct {
	Name              string  `json:"name"`              // 部门名称
	ParentID          string  `json:"parentid"`          // 父部门id，根部门为1
	Order             *string `json:"order"`             // 当前部门在父部门下的所有子部门中的排序值
	CreateDeptGroup   *bool   `json:"createDeptGroup"`   // 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	DeptHiding        *bool   `json:"deptHiding"`        // 是否隐藏部门，true表示隐藏，false表示显示
	DeptPermits       *string `json:"deptPermits"`       // 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的的字符串，使用“|”符号进行分割
	UserPermits       *string `json:"userPermits"`       // 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的的字符串，使用“|”符号进行分割
	OuterDept         *bool   `json:"outerDept"`         // 是否本部门的员工仅可见员工自己，为true时，本部门员工默认只能看到员工自己
	OuterPermitDepts  *string `json:"outerPermitDepts"`  // 本部门的员工仅可见员工自己为true时，可以配置额外可见部门，值为部门id组成的的字符串，使用“|”符号进行分割
	OuterPermitUsers  *string `json:"outerPermitUsers"`  // 本部门的员工仅可见员工自己为true时，可以配置额外可见人员，值为userid组成的的字符串，使用“|”符号进行分割
	OuterDeptOnlySelf *bool   `json:"outerDeptOnlySelf"` // outerDept为true时，可以配置该字段，为true时，表示只能看到所在部门及下级部门通讯录
	SourceIdentifier  *string `json:"sourceIdentifier"`  // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	Ext               *JSON   `json:"ext"`               // 部门自定义字段
}
