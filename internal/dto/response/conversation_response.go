package response

import "github.com/dakong-yi/im-go-server/internal/dto"

type ConversationResponse struct {
	NextSeq          string                   `json:"nextSeq"`
	IsFinished       bool                     `json:"isFinished"`
	ConversationList []*dto.V2TimConversation `json:"conversationList"`
}
