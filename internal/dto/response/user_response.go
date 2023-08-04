package response

import "github.com/dakong-yi/im-go-server/internal/model"

type UserResponse struct {
	UserID        string `json:"userID"`
	NickName      string `json:"nickName"`
	FaceURL       string `json:"faceUrl"`
	SelfSignature string `json:"selfSignature"`
	Gender        int    `json:"gender"`
	AllowType     int    `json:"allowType"`
	Role          int    `json:"role"`
	Level         int    `json:"level"`
	Birthday      int    `json:"birthday"`
}

func ToUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		UserID:        user.UserID,
		NickName:      user.NickName,
		FaceURL:       user.FaceURL,
		SelfSignature: user.SelfSignature,
		Gender:        user.Gender,
		AllowType:     user.AllowType,
		Role:          user.Role,
		Level:         user.Level,
		Birthday:      user.Birthday,
	}
}
