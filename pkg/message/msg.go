package message

// 聊天消息类型
type MessageType int

const (
	MessageTypeText  MessageType = iota // 纯文本消息
	MessageTypeImage                    // 图片消息
	MessageTypeAudio                    // 语音消息
	MessageTypeVideo                    // 视频消息
)

// 聊天形式类型
type ChatType int

const (
	ChatTypePrivate ChatType = iota // 私聊
	ChatTypeGroup                   // 群聊
)

// 消息请求类型
type RequestType int

const (
	RequestTypeChat      RequestType = iota // 聊天
	RequestTypeSignIn                       // 登录or新建连接
	RequestTypeLogOut                       // 退出
	RequestTypeHeartbeat                    // 心跳检测
	RequestTypeResponse                     // 响应结果
)

type MessageRequest struct {
	RequestID string      `json:"request_id"`
	Data      string      `json:"data"`
	Type      RequestType `json:"type"`
}

type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ChatMsg struct {
	Type        MessageType `json:"type"`
	ChatType    ChatType    `json:"chat_type"`
	Sender      string      `json:"sender"`
	Content     string      `json:"content"`
	Target      string      `json:"target"`       // Only used for private messages
	TargetGroup string      `json:"target_group"` // Only used for group messages
}

type SignInMsg struct {
	UserID   string `json:"user_id"`
	DeviceID string `json:"device_id"`
	Token    string `json:"token"`
}
