package controller

import (
	"net/http"
	"strconv"

	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/dakong-yi/im-go-server/internal/dto/response"
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
)

type FriendRequestController struct {
	friendRequestService *service.FriendRequestService
}

func NewFriendRequestController(friendRequestService *service.FriendRequestService) *FriendRequestController {
	return &FriendRequestController{
		friendRequestService: friendRequestService,
	}
}

func (ctrl *FriendRequestController) SendFriendRequest(c *gin.Context) {
	var request model.FriendRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.friendRequestService.SendFriendRequest(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request sent successfully"})
}

func (ctrl *FriendRequestController) AcceptFriendRequest(c *gin.Context) {
	requestIDStr := c.Param("id")
	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	err = ctrl.friendRequestService.AcceptFriendRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request accepted successfully"})
}

func (ctrl *FriendRequestController) RejectFriendRequest(c *gin.Context) {
	requestIDStr := c.Param("id")
	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	err = ctrl.friendRequestService.RejectFriendRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request rejected successfully"})
}

func (ctrl *FriendRequestController) GetPendingFriendRequests(c *gin.Context) {
	userID := c.Param("userID")

	requests, err := ctrl.friendRequestService.GetPendingFriendRequests(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func (ctrl *FriendRequestController) GetFriendApplicationList(c *gin.Context) {
	// userID := c.Param("userID")

	// requests, err := ctrl.friendRequestService.GetFriendRequests(userID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"requests": requests})
	// userIDList := c.Param("userID")
	friendApplicationRes := response.V2TimFriendApplicationResult{
		UnreadCount: 2,
		FriendApplicationList: []*dto.V2TimFriendApplication{
			{
				UserID:     "user2",
				NickName:   "John",
				FaceURL:    "https://www.bugela.com/cjpic/frombd/1/253/1943132031/773911012.jpg",
				AddTime:    1,
				AddSource:  "test",
				AddWording: "test",
				Type:       0,
			},
			{
				UserID:     "user1",
				NickName:   "John",
				FaceURL:    "https://www.bugela.com/cjpic/frombd/1/253/1943132031/773911012.jpg",
				AddTime:    1,
				AddSource:  "test",
				AddWording: "test",
				Type:       1,
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": friendApplicationRes, "code": 0, "desc": "ok"})
}

func (ctrl *FriendRequestController) GetReceivedFriendRequests(c *gin.Context) {
	friendID := c.Param("friendID")

	requests, err := ctrl.friendRequestService.GetReceivedFriendRequests(friendID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func (ctrl *FriendRequestController) CancelFriendRequest(c *gin.Context) {
	requestIDStr := c.Param("id")
	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	err = ctrl.friendRequestService.CancelFriendRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request canceled successfully"})
}
