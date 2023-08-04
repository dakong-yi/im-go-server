package controller

import (
	"net/http"

	"github.com/dakong-yi/im-go-server/internal/dto/response"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
)

type GroupController struct {
	userService *service.UserService
}

func NewGroupController(userService *service.UserService) *GroupController {
	return &GroupController{
		userService: userService,
	}
}

func (c *GroupController) GetJoinedGroupList(ctx *gin.Context) {
	// mock
	res := []response.V2TimGroupInfo{{
		GroupID:         "group123",
		GroupType:       "public",
		GroupName:       "My Group",
		Notification:    "Group notification",
		Introduction:    "Group introduction",
		FaceUrl:         "https://pic3.zhimg.com/v2-24318f30d12288eab6cff92a4ce6afb2_r.jpg",
		IsAllMuted:      true,
		IsSupportTopic:  false,
		Owner:           "user1",
		CreateTime:      1625862378,
		GroupAddOpt:     1,
		LastInfoTime:    1625862378,
		LastMessageTime: 1625862378,
		MemberCount:     10,
		OnlineCount:     5,
		Role:            1,
		RecvOpt:         0,
		JoinTime:        1625862378,
		CustomInfo: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}}

	ctx.JSON(http.StatusOK, gin.H{"data": res, "code": 0, "desc": "success"})
}
