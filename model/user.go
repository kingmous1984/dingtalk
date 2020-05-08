package model

import "time"

// UserDetail  用户详情
type UserDetail struct {
	Unionid         string    `json:"unionid"`         // 员工在当前开发者企业账号范围内的唯一标识，系统生成，固定值，不会改变
	Name            string    `json:"name"`            // 员工名字
	Tel             string    `json:"tel"`             // 分机号（仅限企业内部开发调用）
	WorkPlace       string    `json:"workPlace"`       // 办公地点
	Remark          string    `json:"remark"`          // 备注
	Mobile          string    `json:"mobile"`          // 手机号码
	Email           string    `json:"email"`           // 员工的电子邮箱
	OrgEmail        string    `json:"orgEmail"`        // 员工的企业邮箱，如果员工已经开通了企业邮箱，接口会返回，否则不会返回
	Active          bool      `json:"active"`          // 是否已经激活，true表示已激活，false表示未激活
	OrderInDepts    string    `json:"orderInDepts"`    // 在对应的部门中的排序，Map结构的json字符串，key是部门的Id，value是人员在这个部门的排序值
	IsAdmin         bool      `json:"isAdmin"`         // 是否为企业的老板，true表示是，false表示不是
	IsLeaderInDepts string    `json:"isLeaderInDepts"` // 在对应的部门中是否为主管：Map结构的json字符串，key是部门的Id，value是人员在这个部门中是否为主管，true表示是，false表示不是
	IsHide          bool      `json:"isHide"`          // 是否号码隐藏，true表示隐藏，false表示不隐藏
	Department      []int64   `json:"department"`      // 成员所属部门id列表
	Position        string    `json:"position"`        // 职位信息
	Avatar          string    `json:"avatar"`          // 头像url
	HiredDate       time.Time `json:"hiredDate"`       // 入职时间。Unix时间戳 （在OA后台通讯录中的员工基础信息中维护过入职时间才会返回)
	Jobnumber       string    `json:"jobnumber"`       // 员工工号
	IsSenior        bool      `json:"isSenior"`        // 是否是高管
	StateCode       string    `json:"stateCode"`       // 国家地区码
	Extattr         JSON      `json:"extattr"`         // 扩展属性，可以设置多种属性
	// （手机上最多显示10个扩展属性，具体显示哪些属性，请到OA管理后台->设置->通讯录信息设置和OA管理后台->设置->手机端显示信息设置）。
	// 该字段的值支持链接类型填写，同时链接支持变量通配符自动替换，目前支持通配符有：userid，corpid。
	// 示例： [工位地址](http://www.dingtalk.com?userid=#userid#&corpid=#corpid#)

}

// UserRoles 用户详情角色
type UserRoles struct {
	UserDetail
	Roles []*RoleItem `json:"roles"` // 用户所在角色列表
}

// UserItem 成员信息
type UserItem struct {
	UserID string `json:"userid"` // 员工id
	Name   string `json:"name"`   // 成员名称
}

// UserItemsList 成员列表
type UserItemsList struct {
	HasMore bool        `json:"hasMore"`  // 在分页查询时返回，代表是否还有下一页更多数据
	Items   []*UserItem `json:"userlist"` // 成员列表
}

// UserDetail2  用户详情
type UserDetail2 struct {
	UserID     string    `json:"userid"`     //
	Unionid    string    `json:"unionid"`    // 员工在当前开发者企业账号范围内的唯一标识，系统生成，固定值，不会改变
	Name       string    `json:"name"`       // 员工名字
	Tel        string    `json:"tel"`        // 分机号（仅限企业内部开发调用）
	WorkPlace  string    `json:"workPlace"`  // 办公地点
	Remark     string    `json:"remark"`     // 备注
	Mobile     string    `json:"mobile"`     // 手机号码
	Email      string    `json:"email"`      // 员工的电子邮箱
	OrgEmail   string    `json:"orgEmail"`   // 员工的企业邮箱，如果员工已经开通了企业邮箱，接口会返回，否则不会返回
	Active     bool      `json:"active"`     // 是否已经激活，true表示已激活，false表示未激活
	Order      int64     `json:"order"`      // 表示人员在此部门中的排序，列表是按order的倒序排列输出的，即从大到小排列输出的	（钉钉管理后台里面调整了顺序的话order才有值）
	IsAdmin    bool      `json:"isAdmin"`    // 是否为企业的老板，true表示是，false表示不是
	IsBoss     bool      `json:"isBoss"`     // 是否为企业的老板，true表示是，false表示不是
	IsLeader   bool      `json:"isLeader"`   // 是否是部门的主管，true表示是，false表示不是
	IsHide     bool      `json:"isHide"`     // 是否号码隐藏，true表示隐藏，false表示不隐藏
	Department []int64   `json:"department"` // 成员所属部门id列表
	Position   string    `json:"position"`   // 职位信息
	Avatar     string    `json:"avatar"`     // 头像url
	HiredDate  time.Time `json:"hiredDate"`  // 入职时间。Unix时间戳 （在OA后台通讯录中的员工基础信息中维护过入职时间才会返回)
	Jobnumber  string    `json:"jobnumber"`  // 员工工号
	StateCode  string    `json:"stateCode"`  // 国家地区码
	Extattr    JSON      `json:"extattr"`    // 扩展属性，可以设置多种属性
	// （手机上最多显示10个扩展属性，具体显示哪些属性，请到OA管理后台->设置->通讯录信息设置和OA管理后台->设置->手机端显示信息设置）。
	// 该字段的值支持链接类型填写，同时链接支持变量通配符自动替换，目前支持通配符有：userid，corpid。
	// 示例： [工位地址](http://www.dingtalk.com?userid=#userid#&corpid=#corpid#)

}

// UserDetailList 成员列表
type UserDetailList struct {
	HasMore bool           `json:"hasMore"`  // 在分页查询时返回，代表是否还有下一页更多数据
	Items   []*UserDetail2 `json:"userlist"` // 成员列表
}

// AdminItem 管理员项
type AdminItem struct {
	SysLevel int    `json:"sys_level"` // 管理员角色，1表示主管理员，2表示子管理员
	UserID   string `json:"userid"`    // 员工id
}

// User 用户
type User struct {
	Lang         *string    `json:"lang"`         // 通讯录语言	(默认zh_CN另外支持en_US)
	UserID       *string    `json:"userid"`       // 员工在当前企业内的唯一标识，也称staffId。可由企业在创建时指定，并代表一定含义比如工号，创建后不可修改，企业内必须唯一。	长度为1~64个字符，如果不传，服务器将自动生成一个userid。
	Name         *string    `json:"name"`         // 员工名字
	OrderInDepts *JSON      `json:"orderInDepts"` // 在对应的部门中的排序，	Map结构的json字符串，key是部门的Id, value是人员在这个部门的排序值
	Department   []*int64   `json:"department"`   // 数组类型，数组里面值为整型，成员所属部门id列表
	Position     *string    `json:"position"`     // 职位信息 长度为0~64个字符
	Mobile       *string    `json:"mobile"`       // 手机号码
	Tel          *string    `json:"tel"`          // 分机号（仅限企业内部开发调用）
	WorkPlace    *string    `json:"workPlace"`    // 办公地点 长度为0~50个字符
	Remark       *string    `json:"remark"`       // 备注 长度为0~1000个字符
	Email        *string    `json:"email"`        // 员工的电子邮箱 长度为0~64个字符。企业内必须唯一，不可重复
	OrgEmail     *string    `json:"orgEmail"`     // 员工的企业邮箱，员工的企业邮箱已开通，才能增加此字段， 否则会报错
	Jobnumber    *string    `json:"jobnumber"`    // 员工工号
	IsHide       *bool      `json:"isHide"`       // 是否号码隐藏，true表示隐藏，false表示不隐藏
	IsSenior     *bool      `json:"isSenior"`     // 是否高管模式，	true表示是，false表示不是。
	HiredDate    *time.Time `json:"hiredDate"`    // 入职时间。Unix时间戳 （在OA后台通讯录中的员工基础信息中维护过入职时间才会返回)
	Extattr      *JSON      `json:"extattr"`      // 扩展属性，可以设置多种属性
	// （手机上最多显示10个扩展属性，具体显示哪些属性，请到OA管理后台->设置->通讯录信息设置和OA管理后台->设置->手机端显示信息设置）。
	// 该字段的值支持链接类型填写，同时链接支持变量通配符自动替换，目前支持通配符有：userid，corpid。
	// 示例： [工位地址](http://www.dingtalk.com?userid=#userid#&corpid=#corpid#)

}
