package request

// CreateUserRequest 定义创建用户的请求结构体
type CreateUserRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	NickName      string `json:"nickName"`
	FaceURL       string `json:"faceUrl"`
	SelfSignature string `json:"selfSignature"`
	Gender        int    `json:"gender"`
	AllowType     int    `json:"allowType"`
	Role          string `json:"role"`
	Level         int    `json:"level"`
	Birthday      string `json:"birthday"`
}

// UpdateUserRequest 定义更新用户信息的请求结构体
type UpdateUserRequest struct {
	UserID        string `json:"userID"`
	NickName      string `json:"nickName"`
	FaceURL       string `json:"faceUrl"`
	SelfSignature string `json:"selfSignature"`
	Gender        int    `json:"gender"`
	AllowType     int    `json:"allowType"`
	Role          string `json:"role"`
	Level         int    `json:"level"`
	Birthday      string `json:"birthday"`
}

// LoginRequest 定义用户登录的请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
