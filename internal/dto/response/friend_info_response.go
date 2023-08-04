package response

import "github.com/dakong-yi/im-go-server/internal/dto"

type V2TimFriendInfoResult struct {
	ResultCode int                  `json:"resultCode"`
	ResultInfo string               `json:"resultInfo"`
	Relation   int                  `json:"relation"`
	FriendInfo *dto.V2TimFriendInfo `json:"friendInfo"`
}

type V2TimFriendApplicationResult struct {
	UnreadCount           int                           `json:"unreadCount"`
	FriendApplicationList []*dto.V2TimFriendApplication `json:"friendApplicationList"`
}
