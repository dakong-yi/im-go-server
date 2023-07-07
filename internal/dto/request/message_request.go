package request

type CreateMessageRequest struct {
	ConversationID int    `json:"conversation_id" mapstructure:"conversation_id"`
	SenderID       string `json:"sender_id" mapstructure:"sender_id"`
	Content        string `json:"content" mapstructure:"content"`
	Type           string `json:"type" mapstructure:"type"`
	ImageURL       string `json:"image_url" mapstructure:"image_url"`
	AudioURL       string `json:"audio_url" mapstructure:"audio_url"`
	VideoURL       string `json:"video_url" mapstructure:"video_url"`
	FileURL        string `json:"file_url" mapstructure:"file_url"`
	// 添加其他字段...
}
