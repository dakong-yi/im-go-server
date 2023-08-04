package model

import (
	"gorm.io/gorm"
)

const FriendRequestStatusAccepted = "accepted"
const FriendRequestStatusRejected = "rejected"
const FriendRequestStatusPending = "pending"

type FriendRequest struct {
	gorm.Model
	UserID   string // 用户ID
	FriendID string // 好友ID
	Message  string // 请求消息
	Status   string // 请求状态：pending, accepted, rejected
}
