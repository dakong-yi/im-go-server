package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        string `json:"user_id"`
	Identifier    string `json:"identifier"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	NickName      string `json:"nick_name"`
	FaceURL       string `json:"face_url"`
	SelfSignature string `json:"self_signature"`
	Gender        int    `json:"gender"`
	AllowType     int    `json:"allow_type"`
	Role          string `json:"role"`
	Level         int    `json:"level"`
	Birthday      string `json:"birthday"`
	// 其他字段...
}
