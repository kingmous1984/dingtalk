package model

import (
	"encoding/json"
)

// MessageType 消息类型
type MessageType string

const (
	// MessageTypeText 文本消息
	MessageTypeText MessageType = "text"
	// MessageTypeImage 图片消息
	MessageTypeImage MessageType = "image"
	// MessageTypeVoice 语音消息
	MessageTypeVoice MessageType = "voice"
	// MessageTypeFile 文件消息
	MessageTypeFile MessageType = "file"
	// MessageTypeLink 链接消息
	MessageTypeLink MessageType = "link"
	// MessageTypeOA OA消息
	MessageTypeOA MessageType = "oa"
	// MessageTypeMarkdown 说明：目前仅支持以下md语法; markdown格式的消息，建议500字符以内
	MessageTypeMarkdown MessageType = "markdown"
	// MessageTypeActionCard 卡片消息 卡片消息支持整体跳转ActionCard样式和独立跳转ActionCard样式：
	// （1）整体跳转ActionCard样式，支持一个点击Action，需要传入参数 single_title和 single_url；
	// （2）独立跳转ActionCard样式，支持多个点击Action，需要传入参数 btn_orientation 和 btn_json_list；
	MessageTypeActionCard MessageType = "action_card"
)

// ChatCreate 创建会话结构
type ChatCreate struct {
	Name                string   `json:"name"`                // 群名称，长度限制为1~20个字符
	Owner               string   `json:"owner"`               // 群主userId，员工唯一标识ID；必须为该会话useridlist的成员之一
	Useridlist          []string `json:"useridlist"`          // 群成员列表，每次最多支持40人，群人数上限为1000
	ShowHistoryType     *int     `json:"showHistoryType"`     // 新成员是否可查看聊天历史消息（新成员入群是否可查看最近100条聊天记录），0代表否；1代表是；不传默认为否
	Searchable          *int     `json:"searchable"`          // 群可搜索，0-默认，不可搜索，1-可搜索
	ValidationType      *int     `json:"validationType"`      // 入群验证，0：不入群验证（默认） 1：入群验证
	MentionAllAuthority *int     `json:"mentionAllAuthority"` // @ all 权限，0-默认，所有人，1-仅群主可@ all
	ChatBannedType      *int     `json:"chatBannedType"`      // 群禁言，0-默认，不禁言，1-全员禁言
	ManagementType      *int     `json:"managementType"`      // 管理类型，0-默认，所有人可管理，1-仅群主可管理
}

// ChatUpdate 修改会话结构
type ChatUpdate struct {
	ChatID              string    `json:"chatid"`              // 群会话的id。仅支持通过调用服务端“创建会话”接口获取的chatid，不支持通过调用前端JSAPI获取的chatid。
	Name                *string   `json:"name"`                // 群名称，长度限制为1~20个字符
	Owner               *string   `json:"owner"`               // 群主userId，员工唯一标识ID；必须为该会话useridlist的成员之一
	AddUserIDList       []*string `json:"add_useridlist"`      // 添加成员列表，每次最多支持40人，群人数上限为1000
	DelUserIDList       []*string `json:"del_useridlist"`      // 删除成员列表，每次最多支持40人，群人数上限为1000
	Icon                *string   `json:"icon"`                // 群头像mediaid
	ChatBannedType      *int      `json:"chatBannedType"`      // 群禁言，0-默认，不禁言，1-全员禁言
	Searchable          *int      `json:"searchable"`          // 群可搜索，0-默认，不可搜索，1-可搜索
	ShowHistoryType     *int      `json:"showHistoryType"`     // 新成员是否可查看聊天历史消息（新成员入群是否可查看最近100条聊天记录），0代表否；1代表是；不传默认为否
	ValidationType      *int      `json:"validationType"`      // 入群验证，0：不入群验证（默认） 1：入群验证
	MentionAllAuthority *int      `json:"mentionAllAuthority"` // @ all 权限，0-默认，所有人，1-仅群主可@ all
	ManagementType      *int      `json:"managementType"`      // 管理类型，0-默认，所有人可管理，1-仅群主可管理
}

// ChatInfo 会话信息结构
type ChatInfo struct {
	ChatID              string   `json:"chatid"`              // 群会话的id
	Name                string   `json:"name"`                // 群名称，长度限制为1~20个字符
	Owner               string   `json:"owner"`               // 群主userId，员工唯一标识ID；必须为该会话useridlist的成员之一
	Useridlist          []string `json:"useridlist"`          // 群成员列表，每次最多支持40人，群人数上限为1000
	ShowHistoryType     int      `json:"showHistoryType"`     // 新成员是否可查看聊天历史消息（新成员入群是否可查看最近100条聊天记录），0代表否；1代表是；不传默认为否
	Searchable          int      `json:"searchable"`          // 群可搜索，0-默认，不可搜索，1-可搜索
	ValidationType      int      `json:"validationType"`      // 入群验证，0：不入群验证（默认） 1：入群验证
	MentionAllAuthority int      `json:"mentionAllAuthority"` // @ all 权限，0-默认，所有人，1-仅群主可@ all
	ChatBannedType      int      `json:"chatBannedType"`      // 群禁言，0-默认，不禁言，1-全员禁言
	ManagementType      int      `json:"managementType"`      // 管理类型，0-默认，所有人可管理，1-仅群主可管理
	Icon                string   `json:"icon"`                // 群头像mediaid
}

// GetTextMessage 获取text 消息类型的json
//----------------------------------------------------//
func GetTextMessage(info string) ([]byte, error) {
	data := struct {
		MsgType MessageType `json:"msgtype"` // 消息类型，此时固定为：text
		Text    struct {
			Content string `json:"content"` // 消息内容，建议500字符以内
		} `json:"text"`
	}{
		MsgType: MessageTypeText,
	}
	data.Text.Content = info
	return json.Marshal(data)
}

// GetImageMessage 获取image 消息类型的json
//----------------------------------------------------//
func GetImageMessage(mediaID string) ([]byte, error) {
	data := struct {
		MsgType MessageType `json:"msgtype"` // 消息类型，此时固定为：image
		Image   struct {
			MediaID string `json:"media_id"` // 媒体文件id。可以通过媒体文件接口上传图片获取。建议宽600像素 x 400像素，宽高比3 : 2
		} `json:"image"`
	}{
		MsgType: MessageTypeImage,
	}
	data.Image.MediaID = mediaID
	return json.Marshal(data)
}

// GetVoiceMessage 获取voice 消息类型的json
//----------------------------------------------------//
func GetVoiceMessage(mediaID string, duration string) ([]byte, error) {
	data := struct {
		MsgType MessageType `json:"msgtype"` // 消息类型，此时固定为：voice
		Voice   struct {
			MediaID  string `json:"media_id"` // 媒体文件id。2MB，播放长度不超过60s，AMR格式。可以通过媒体文件接口上传图片获取。
			Duration string `json:"duration"` // 正整数，小于60，表示音频时长
		} `json:"voice"`
	}{
		MsgType: MessageTypeVoice,
	}
	data.Voice.MediaID = mediaID
	data.Voice.Duration = duration
	return json.Marshal(data)
}

// GetFileMessage 获取file 消息类型的json
//----------------------------------------------------//
func GetFileMessage(mediaID string) ([]byte, error) {
	data := struct {
		MsgType MessageType `json:"msgtype"` // 消息类型，此时固定为：file
		File    struct {
			MediaID string `json:"media_id"` // 媒体文件id。引用的媒体文件最大10MB。可以通过媒体文件接口上传图片获取。
		} `json:"file"`
	}{
		MsgType: MessageTypeFile,
	}
	data.File.MediaID = mediaID
	return json.Marshal(data)
}

// Link 连接消息结构
//----------------------------------------------------//
type Link struct {
	MessageURL string `json:"messageUrl"` // 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接
	PicURL     string `json:"picUrl"`     // 图片地址。可以通过媒体文件接口上传图片获取。
	Title      string `json:"title"`      // 消息标题，建议100字符以内
	Text       string `json:"text"`       // 消息描述，建议500字符以内
}

// GetLinkMessage 获取link 消息类型的json
func GetLinkMessage(link *Link) ([]byte, error) {
	data := struct {
		MsgType MessageType `json:"msgtype"` // 消息类型，此时固定为：link
		MsgLink *Link       `json:"link"`
	}{
		MsgType: MessageTypeLink,
		MsgLink: link,
	}
	return json.Marshal(data)
}

// OA oa消息结构
//----------------------------------------------------//
type OA struct {
	MessageURL   string  `json:"message_url"`    // 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接
	PcMessageURL *string `json:"pc_message_url"` // PC端点击消息时跳转到的地址
	Head         OAHead  `json:"head"`           // 消息头部内容
	Body         OABody  `json:"body"`           // 消息体
}

// OAHead 消息头部内容
type OAHead struct {
	BgColor string `json:"bgcolor"` // 消息头部的背景颜色。长度限制为8个英文字符，其中前2为表示透明度，后6位表示颜色值。不要添加0x
	Text    string `json:"text"`    // 消息的头部标题 (向普通会话发送时有效，向企业会话发送时会被替换为微应用的名字)，长度限制为最多10个字符
}

// OABody 消息体
type OABody struct {
	Title     *string       `json:"title"`      // 消息体的标题，建议50个字符以内
	Form      []*OABodyForm `json:"form"`       // 消息体的表单，最多显示6个，超过会被隐藏
	Rich      OABodyRich    `json:"rich"`       // 单行富文本信息
	Content   *string       `json:"content"`    // 消息体的内容，最多显示3行
	Image     *string       `json:"image"`      // 消息体中的图片，支持图片资源@mediaId
	FileCount *string       `json:"file_count"` // 自定义的附件数目。此数字仅供显示，钉钉不作验证
	Author    *string       `json:"author"`     // 自定义的作者名字
}

// OABodyForm 消息体的表单
type OABodyForm struct {
	Key   *string `json:"key"`   // 消息体的关键字
	Value *string `json:"value"` // 消息体的关键字对应的值
}

// OABodyRich 单行富文本信息
type OABodyRich struct {
	Num  *string `json:"num"`  // 单行富文本信息的数目
	Unit *string `json:"unit"` // 单行富文本信息的单位
}

// GetOAMessage 获取oa 消息类型的json
func GetOAMessage(oa *OA) ([]byte, error) {
	data := struct {
		MsgType MessageType `json:"msgtype"` // 消息类型，此时固定为：oa
		MsgOA   *OA         `json:"oa"`
	}{
		MsgType: MessageTypeOA,
		MsgOA:   oa,
	}
	return json.Marshal(data)
}

// Markdown 结构
//----------------------------------------------------//
type Markdown struct {
	Title string `json:"title"` // 首屏会话透出的展示内容
	Text  string `json:"text"`  // markdown格式的消息，建议500字符以内
}

// GetMarkdownMessage 获取Markdown 消息类型的json
func GetMarkdownMessage(markdown *Markdown) ([]byte, error) {
	data := struct {
		MsgType  MessageType `json:"msgtype"` // 消息类型，此时固定为：markdown
		Markdown *Markdown   `json:"markdown"`
	}{
		MsgType:  MessageTypeLink,
		Markdown: markdown,
	}
	return json.Marshal(data)
}

// ActionCard 卡片消息
//----------------------------------------------------//
type ActionCard struct {
	Title          string     `json:"title"`           // 透出到会话列表和通知的文案，最长64个字符
	Markdown       string     `json:"markdown"`        // 消息内容，支持markdown，语法参考标准markdown语法。建议1000个字符以内
	SingleTitle    *string    `json:"single_title"`    // 使用整体跳转ActionCard样式时的标题，必须与single_url同时设置，最长20个字符
	SingleURL      *string    `json:"single_url"`      // 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接，最长500个字符
	BtnOrientation *string    `json:"btn_orientation"` // 使用独立跳转ActionCard样式时的按钮排列方式，竖直排列(0)，横向排列(1)；必须与btn_json_list同时设置
	BtnJSONList    []*BtnJSON `json:"btn_json_list"`   // 使用独立跳转ActionCard样式时的按钮列表；必须与btn_orientation同时设置
}

// BtnJSON 使用独立跳转ActionCard样式时的按钮
type BtnJSON struct {
	Title     *string `json:"title"`      // 使用独立跳转ActionCard样式时的按钮的标题，最长20个字符
	ActionURL *string `json:"action_url"` // 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接，最长500个字符
}

// GetActionCardMessage 获取ActionCard 消息类型的json
func GetActionCardMessage(actionCard *ActionCard) ([]byte, error) {
	data := struct {
		MsgType    MessageType `json:"msgtype"` // 消息类型，此时固定为：markdown
		ActionCard *ActionCard `json:"action_card"`
	}{
		MsgType:    MessageTypeLink,
		ActionCard: actionCard,
	}
	return json.Marshal(data)
}
