package request

type CreateConversationRequest struct {
	Name             string   `json:"name" binding:"required"`
	OwnerID          string   `json:"owner_id" binding:"required"`
	UserIDs          []string `json:"user_ids" binding:"required"`
	ConversationType string   `json:"type" binding:"required"`
}
