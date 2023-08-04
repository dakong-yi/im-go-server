package controller

import (
	"net/http"

	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/dto/response"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
)

type FriendshipController struct {
	friendshipService *service.FriendshipService
}

func NewFriendshipController(friendshipService *service.FriendshipService) *FriendshipController {
	return &FriendshipController{
		friendshipService: friendshipService,
	}
}

func (ctrl *FriendshipController) AddFriend(c *gin.Context) {
	userID := c.Param("userID")
	friendID := c.Param("friendID")

	err := ctrl.friendshipService.CreateFriendship(userID, friendID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend added successfully"})
}

func (ctrl *FriendshipController) RemoveFriend(c *gin.Context) {
	userID := c.Param("userID")
	friendID := c.Param("friendID")

	err := ctrl.friendshipService.DeleteFriendship(userID, friendID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend removed successfully"})
}

func (ctrl *FriendshipController) GetFriends(c *gin.Context) {
	userID := c.Param("userID")

	friends, err := ctrl.friendshipService.GetFriends(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": friends, "code": 0, "desc": "ok"})
}

// 更新好友备注
func (c *FriendshipController) UpdateFriendRemark(ctx *gin.Context) {
	userID := ctx.Param("userID")
	friendID := ctx.Param("friendID")
	remark := ctx.PostForm("remark")

	err := c.friendshipService.UpdateFriendRemark(userID, friendID, remark)
	if err != nil {
		// 处理错误
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update friend remark",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Friend remark updated successfully",
	})
}

// 获取好友备注
func (c *FriendshipController) GetFriendRemark(ctx *gin.Context) {
	userID := ctx.Param("userID")
	friendID := ctx.Param("friendID")

	remark, err := c.friendshipService.GetFriendRemark(userID, friendID)
	if err != nil {
		// 处理错误
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get friend remark",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"remark": remark,
	})
}

func (ctrl *FriendshipController) GetFriendsInfo(c *gin.Context) {
	var request request.GetFriendsInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	friendsInfo, err := ctrl.friendshipService.GetFriendsInfo(request.Sender, request.UserIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := []response.V2TimFriendInfoResult{}
	for _, friend := range friendsInfo {
		res = append(res, response.V2TimFriendInfoResult{
			ResultCode: 200,
			ResultInfo: "",
			Relation:   3,
			FriendInfo: friend,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": res, "code": 0, "desc": "ok"})
}

// TODO
func (ctrl *FriendshipController) CheckFriend(c *gin.Context) {
	// userIDList := c.Param("userID")
	res := []response.V2TimFriendCheckResult{
		{
			UserID:     "123",
			ResultCode: 200,
			ResultInfo: "Success",
			ResultType: 3,
		},
	}
	c.JSON(http.StatusOK, gin.H{"data": res, "code": 0, "desc": "ok"})
}
