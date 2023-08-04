package dto

type V2TimUserFullInfo struct {
	UserID        string `json:"userID"`
	NickName      string `json:"nickName"`
	FaceURL       string `json:"faceUrl"`
	SelfSignature string `json:"selfSignature"`
	Gender        int    `json:"gender"`
	AllowType     int    `json:"allowType"`
	Role          int    `json:"role"`
	Level         int    `json:"level"`
	Birthday      int    `json:"birthday"`
}

type V2TimUserStatus struct {
	UserID       string `json:"userID"`
	StatusType   int    `json:"statusType"`
	CustomStatus string `json:"customStatus"`
}
