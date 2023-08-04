package dto

type V2TimFriendInfo struct {
	UserID           string            `json:"userID"`
	FriendRemark     string            `json:"friendRemark"`
	FriendGroups     []string          `json:"friendGroups"`
	FriendCustomInfo map[string]string `json:"friendCustomInfo"`
	UserProfile      V2TimUserFullInfo `json:"userProfile"`
}

type V2TimFriendApplication struct {
	UserID     string `json:"userID"`
	NickName   string `json:"nickName"`
	FaceURL    string `json:"faceUrl"`
	AddTime    int64  `json:"addTime"`    //添加时间
	AddSource  string `json:"addSource"`  //来源
	AddWording string `json:"addWording"` //加好友附言
	Type       int    `json:"type"`       //0:别人发给我的1:我发给别人的2:别人发给我的 和 我发给别人的。仅拉取时有效
}
