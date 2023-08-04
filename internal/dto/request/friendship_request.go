package request

type GetFriendsInfoRequest struct {
	UserIDs []string `json:"user_ids"`
	Sender  string   `json:"sender"`
}
