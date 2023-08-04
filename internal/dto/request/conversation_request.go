package request

type CreateConversationRequest struct {
	OwnerID string `json:"owner_id" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
	GroupID string `json:"group_id" binding:"required"`
}
