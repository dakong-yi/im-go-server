package controller

import (
	"net/http"

	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
)

type BlacklistController struct {
	blacklistService *service.BlacklistService
}

func NewBlacklistController(blacklistService *service.BlacklistService) *BlacklistController {
	return &BlacklistController{
		blacklistService: blacklistService,
	}
}

func (c *BlacklistController) AddToBlacklist(ctx *gin.Context) {
	userID := ctx.Param("userID")
	blacklistUserID := ctx.Param("blacklistUserID")

	err := c.blacklistService.AddToBlacklist(userID, blacklistUserID)
	if err != nil {
		// 处理错误
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add user to blacklist",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User added to blacklist successfully",
	})
}

func (c *BlacklistController) RemoveFromBlacklist(ctx *gin.Context) {
	userID := ctx.Param("userID")
	blacklistUserID := ctx.Param("blacklistUserID")

	err := c.blacklistService.RemoveFromBlacklist(userID, blacklistUserID)
	if err != nil {
		// 处理错误
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to remove user from blacklist",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User removed from blacklist successfully",
	})
}

func (c *BlacklistController) GetBlacklist(ctx *gin.Context) {
	// userID := ctx.Param("userID")

	// blacklist, err := c.blacklistService.GetBlacklistByUserID(userID)
	// if err != nil {
	// 	// 处理错误
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to get blacklist",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"blacklist": blacklist,
	// })
	res := []dto.V2TimFriendInfo{
		{
			UserID:           "user3",
			FriendRemark:     "user3",
			FriendGroups:     []string{},
			FriendCustomInfo: map[string]string{},
			UserProfile: dto.V2TimUserFullInfo{
				NickName:      "user3",
				UserID:        "user3",
				FaceURL:       "https://pic3.zhimg.com/v2-af5b7b4328b26e35c4b1860b1baad31e_r.jpg",
				SelfSignature: "user3",
				Gender:        1,
				AllowType:     1,
				Role:          2,
				Level:         5,
				Birthday:      19900101,
			},
		},
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res, "code": 0, "desc": "ok"})
}
