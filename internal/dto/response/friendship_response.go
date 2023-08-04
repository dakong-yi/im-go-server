package response

type V2TimFriendCheckResult struct {
	UserID     string `json:"userID"`
	ResultCode int    `json:"resultCode"`
	ResultInfo string `json:"resultInfo"`
	ResultType int    `json:"resultType"`
}
