package response

import "github.com/dakong-yi/im-go-server/internal/dto"

type V2TimMessageListResult struct {
	IsFinished  bool                `json:"isFinished"`
	MessageList []*dto.V2TimMessage `json:"messageList"`
}

type V2TimMsgCreateInfoResult struct {
	ID          string            `json:"id"`
	MessageInfo *dto.V2TimMessage `json:"messageInfo"`
}
