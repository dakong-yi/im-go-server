package response

import (
	"time"

	"github.com/dakong-yi/im-go-server/internal/model"
)

type MessageResponse struct {
	ID             int               `json:"id" mapstructure:"id"`
	ConversationID int               `json:"conversation_id" mapstructure:"conversation_id"`
	SenderID       string            `json:"sender_id" mapstructure:"sender_id"`
	Content        string            `json:"content" mapstructure:"content"`
	Type           model.MessageType `json:"type" mapstructure:"type"`
	ImageURL       string            `json:"image_url" mapstructure:"image_url"`
	AudioURL       string            `json:"audio_url" mapstructure:"audio_url"`
	VideoURL       string            `json:"video_url" mapstructure:"video_url"`
	FileURL        string            `json:"file_url" mapstructure:"file_url"`
	Timestamp      time.Time         `json:"timestamp" mapstructure:"timestamp"`
}

func ToMessageResponse(msg *model.Message) *MessageResponse {
	message := &MessageResponse{
		ID:             int(msg.ID),
		ConversationID: msg.ConversationID,
		SenderID:       msg.SenderID,
		Content:        msg.Content,
		Type:           msg.Type,
		ImageURL:       msg.ImageURL,
		AudioURL:       msg.AudioURL,
		VideoURL:       msg.VideoURL,
		FileURL:        msg.FileURL,
		Timestamp:      msg.Timestamp,
	}
	return message
}
