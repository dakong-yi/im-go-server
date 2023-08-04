package response

type V2TimGroupInfo struct {
	GroupID         string            `json:"groupID"`
	GroupType       string            `json:"groupType"`
	GroupName       string            `json:"groupName"`
	Notification    string            `json:"notification"`
	Introduction    string            `json:"introduction"`
	FaceUrl         string            `json:"faceUrl"`
	IsAllMuted      bool              `json:"isAllMuted"`
	IsSupportTopic  bool              `json:"isSupportTopic"`
	Owner           string            `json:"owner"`
	CreateTime      int               `json:"createTime"`
	GroupAddOpt     int               `json:"groupAddOpt"`
	LastInfoTime    int               `json:"lastInfoTime"`
	LastMessageTime int               `json:"lastMessageTime"`
	MemberCount     int               `json:"memberCount"`
	OnlineCount     int               `json:"onlineCount"`
	Role            int               `json:"role"`
	RecvOpt         int               `json:"recvOpt"`
	JoinTime        int               `json:"joinTime"`
	CustomInfo      map[string]string `json:"customInfo"`
}
