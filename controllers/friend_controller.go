// controllers/friend_controller.go

package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/dakong-yi/im-go-server/models"
// 	"github.com/dakong-yi/im-go-server/services"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pkg/errors"
// )

// type FriendController struct {
// 	friendService *services.FriendService
// }

// func NewFriendController(friendService *services.FriendService) *FriendController {
// 	return &FriendController{friendService: friendService}
// }

// func (c *FriendController) CreateFriend(ctx *gin.Context) {
// 	var friend models.Friend
// 	err := ctx.ShouldBindJSON(&friend)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
// 		return
// 	}

// 	err = c.friendService.Create(&friend)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to create friend").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "friend created successfully"})
// }

// func (c *FriendController) DeleteFriend(ctx *gin.Context) {
// 	friendID, err := strconv.ParseUint(ctx.Param("friend_id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid friend ID"})
// 		return
// 	}

// 	friend, err := c.friendService.FindByID(uint(friendID))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to find friend").Error()})
// 		return
// 	}

// 	if friend == nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "friend not found"})
// 		return
// 	}

// 	err = c.friendService.Delete(friend)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to delete friend").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "friend deleted successfully"})
// }

// func (c *FriendController) GetFriendsByUserID(ctx *gin.Context) {
// 	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
// 		return
// 	}

// 	friendIDs, err := c.friendService.FindFriendsByUserID(uint(userID))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to find friends").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"friend_ids": friendIDs})
// }

// func (c *FriendController) GetFriendsByFriendID(ctx *gin.Context) {
// 	friendID, err := strconv.ParseUint(ctx.Param("friend_id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid friend ID"})
// 		return
// 	}

// 	userIDs, err := c.friendService.FindFriendsByFriendID(uint(friendID))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to find friends").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"user_ids": userIDs})
// }

// func (c *FriendController) GetUsersByFriendID(ctx *gin.Context) {
// 	friendID, err := strconv.ParseUint(ctx.Param("friend_id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid friend ID"})
// 		return
// 	}

// 	users, err := c.friendService.FindUsersByFriendID(uint(friendID))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to find users").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"users": users})
// }
