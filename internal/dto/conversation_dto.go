package dto

type V2TimConversation struct {
	ConversationID        string        `json:"conversationID"`
	Type                  int           `json:"type,omitempty"`
	UserID                string        `json:"userID,omitempty"`
	GroupID               string        `json:"groupID,omitempty"`
	ShowName              string        `json:"showName,omitempty"`
	FaceUrl               string        `json:"faceUrl,omitempty"`
	GroupType             string        `json:"groupType,omitempty"`
	UnreadCount           int           `json:"unreadCount,omitempty"`
	LastMessage           *V2TimMessage `json:"lastMessage,omitempty"`
	DraftText             string        `json:"draftText,omitempty"`
	DraftTimestamp        int64         `json:"draftTimestamp,omitempty"`
	GroupAtInfoList       []interface{} `json:"groupAtInfoList,omitempty"`
	IsPinned              bool          `json:"isPinned"`
	RecvOpt               int           `json:"recvOpt,omitempty"`
	Orderkey              int64         `json:"orderkey,omitempty"`
	MarkList              []int         `json:"markList,omitempty"`
	CustomData            string        `json:"customData,omitempty"`
	ConversationGroupList []string      `json:"conversationGroupList,omitempty"`
}
