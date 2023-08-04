package controller

import (
	"net/http"

	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	// 从请求中获取用户注册信息
	var user request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用用户服务层的注册方法
	if _, err := c.userService.Register(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ctrl *UserController) Login(c *gin.Context) {
	//mock
	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "Login successfully"})
	return
	// 解析请求参数
	var request request.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 调用用户服务层进行登录
	user, err := ctrl.userService.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUserInfo(c *gin.Context) {
	//mock
	// 解析请求参数
	var request request.GetUserInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 调用用户服务层进行登录
	user, err := ctrl.userService.GetUserInfo(request.UserIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "", "data": user})
}

func (ctrl *UserController) GetUserStatus(c *gin.Context) {
	//mock
	res := []dto.V2TimUserStatus{
		{
			UserID:     "user1",
			StatusType: 1,
		},
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "", "data": res})
}
