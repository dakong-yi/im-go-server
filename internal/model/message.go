package model

import (
	"time"

	"gorm.io/gorm"
)

type ElemType int

const (
	UNKNOWN_MESSAGE    ElemType = iota // 0
	TEXT_MESSAGE                       // 1
	CUSTOM_MESSAGE                     // 2
	IMAGE_MESSAGE                      // 3
	VOICE_MESSAGE                      // 4
	VIDEO_MESSAGE                      // 5
	FILE_MESSAGE                       // 6
	LOCATION_MESSAGE                   // 7
	EMOJI_MESSAGE                      // 8
	GROUP_TIPS_MESSAGE                 // 9
	MERGED_MESSAGE                     // 10
)

type MessageStatus int

const (
	MESSAGE_SENDING_IN_PROGRESS MessageStatus = iota + 1
	MESSAGE_SENT_SUCCESSFULLY
	MESSAGE_SENDING_FAILED
	MESSAGE_DELETED
	MESSAGE_IMPORTED_LOCALLY
	MESSAGE_REVOKED
)

const (
	V2TIM_PRIORITY_DEFAULT int = iota
	V2TIM_PRIORITY_HIGH
	V2TIM_PRIORITY_NORMAL
	V2TIM_PRIORITY_LOW
)

var PRIORITY = map[string]int{
	"V2TIM_PRIORITY_DEFAULT": 0,
	"V2TIM_PRIORITY_HIGH":    1,
	"V2TIM_PRIORITY_NORMAL":  2,
	"V2TIM_PRIORITY_LOW":     3,
}

const (
	V2TIM_NULL int = iota
	V2TIM_GET_CLOUD_OLDER_MSG
	V2TIM_GET_CLOUD_NEWER_MSG
	V2TIM_GET_LOCAL_OLDER_MSG
	V2TIM_GET_LOCAL_NEWER_MSG
)

type Message struct {
	MsgID        string        `json:"msg_id"`
	Timestamp    int64         `json:"timestamp"`
	Progress     int           `json:"progress"`
	Sender       string        `json:"sender"`
	NickName     string        `json:"nick_name"`
	FriendRemark string        `json:"friend_remark"`
	FaceUrl      string        `json:"face_url"`
	NameCard     string        `json:"name_card"`
	GroupID      string        `json:"group_id"`
	UserID       string        `json:"user_id"`
	Status       MessageStatus `json:"status"`
	ElemType     ElemType      `json:"elem_type"`
	IsSelf       bool          `json:"is_self"`
	IsRead       bool          `json:"is_read"`
	IsPeerRead   bool          `json:"is_peer_read"` //消息对方是否已读（只有 C2C 消息有效） 该字段为true的条件是消息 timestamp <= 对端标记会话已读的时间
	Priority     int           `json:"priority"`
	// GroupAtUserList           []string      `json:"group_at_user_list"`
	Seq                       string `json:"seq"` // 群聊中的消息序列号云端生成，在群里是严格递增且唯一的, 单聊中的序列号是本地生成，不能保证严格递增且唯一。
	Random                    int    `json:"random"`
	IsExcludedFromUnreadCount bool   `json:"is_excluded_from_unread_count"` // 消息是否不计入会话未读数
	IsExcludedFromLastMessage bool   `json:"is_excluded_from_last_message"` // 消息是否不计入会话 lastMsg
	ID                        string `json:"id"`
	NeedReadReceipt           bool   `json:"need_read_receipt"`
	Text                      string `json:"text"`
	CloudCustomData           string `json:"cloud_custom_data"`
	LocalCustomData           string `json:"local_custom_data"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	DeletedAt                 gorm.DeletedAt `gorm:"index"`
}
