package request

type CreateTextMessageRequest struct {
	Sender   string `json:"sender"`
	Text     string `json:"text"`
	ElemType int    `json:"elemType"`
}
type SendMessageRequest struct {
	ID             string `json:"id"`
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	GroupID        string `json:"groupID"`
	Priority       string `json:"priority"`
	OnlineUserOnly bool   `json:"onlineUserOnly"`
	// OfflinePushInfo           *dto.OfflinePushInfo `json:"offlinePushInfo"`
	NeedReadReceipt           bool   `json:"needReadReceipt"`
	IsExcludedFromUnreadCount bool   `json:"isExcludedFromUnreadCount"`
	IsExcludedFromLastMessage bool   `json:"isExcludedFromLastMessage"`
	CloudCustomData           string `json:"cloudCustomData"`
	// LocalCustomData           string `json:"localCustomData"`
}

type GetHistoryMessageListRequest struct {
	Count           int    `json:"count"`
	GetType         int    `json:"getType"`
	Sender          string `json:"sender"`
	UserID          string `json:"userID"`
	GroupID         string `json:"groupID"`
	LastMsgID       string `json:"lastMsgID"`
	LastMsgSeq      int    `json:"lastMsgSeq"`
	MessageTypeList []int  `json:"messageTypeList"`
}
type MarkC2CMessageAsReadRequest struct {
	OwnerID string `json:"ownerID"`
	UserID  string `json:"userID"`
}
