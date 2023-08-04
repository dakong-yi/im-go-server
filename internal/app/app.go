package app

import (
	"log"

	"github.com/dakong-yi/im-go-server/internal/controller"
	"github.com/dakong-yi/im-go-server/internal/http"
	"github.com/dakong-yi/im-go-server/internal/repository"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/dakong-yi/im-go-server/internal/socket"
	"go.uber.org/dig"
)

func Run() {
	container := dig.New()

	container.Provide(http.NewServer)
	container.Provide(socket.NewSocketServer)
	container.Provide(controller.NewConversationController)
	container.Provide(controller.NewMessageController)
	container.Provide(controller.NewUserController)
	container.Provide(service.NewUserService)
	container.Provide(repository.NewUserRepositoryImpl)
	container.Provide(service.NewMessageService)
	container.Provide(service.NewConversationService)
	container.Provide(repository.NewMessageRepoImpl)
	container.Provide(repository.NewConversationRepoImpl)
	container.Provide(repository.NewUserConversationRepoImpl)
	container.Provide(repository.NewFriendshipRepoImpl)
	container.Provide(repository.NewFriendRequestRepoImpl)
	container.Provide(repository.NewBlacklistRepoImpl)
	container.Provide(service.NewBlacklistService)
	container.Provide(service.NewFriendshipService)
	container.Provide(service.NewFriendRequestService)
	container.Provide(controller.NewBlacklistController)
	container.Provide(controller.NewFriendshipController)
	container.Provide(controller.NewFriendRequestController)
	container.Provide(controller.NewGroupController)
	// 增加其他依赖...

	err := container.Invoke(func(server *http.Server, socket *socket.SocketServer,
		userController *controller.UserController,
		conversationController *controller.ConversationController,
		messageController *controller.MessageController,
		blacklistController *controller.BlacklistController,
		friendshipController *controller.FriendshipController,
		friendRequestController *controller.FriendRequestController,
		GroupController *controller.GroupController,
		// 增加其他依赖...
	) {
		// 在这里可以使用注入后的控制器注册路由
		server.Router.POST("/login", userController.Login)
		server.Router.POST("/register", userController.Register)
		server.Router.POST("/users/info", userController.GetUserInfo)

		server.Router.POST("/conversation", conversationController.CreateConversation)
		server.Router.GET("/conversation/:ownerID/:conversationID", conversationController.GetConversation)
		server.Router.GET("/conversation/users/:userID", conversationController.GetConversationsByUserID)
		server.Router.GET("/conversation/totalUnread/:ownerID", conversationController.GetTotalUnreadMessageCount)

		// 注册路由和处理函数
		server.Router.POST("/message", messageController.CreateMessage)
		server.Router.POST("/message/send", messageController.SendMessage)
		server.Router.POST("/message/history", messageController.GetHistoryMessageList)
		server.Router.POST("/message/history/v2", messageController.GetHistoryMessageListV2)
		server.Router.GET("/message/:id", messageController.GetMessageByMsgID)
		server.Router.POST("/message/markC2CMessageAsRead", messageController.MarkC2CMessageAsRead)

		// 添加用户到黑名单
		server.Router.POST("/users/:userID/blacklist/:blacklistUserID", blacklistController.AddToBlacklist)
		// 从黑名单移除用户
		server.Router.DELETE("/users/:userID/blacklist/:blacklistUserID", blacklistController.RemoveFromBlacklist)
		// 获取黑名单列表
		server.Router.GET("/users/:userID/blacklist", blacklistController.GetBlacklist)

		// 发送好友请求
		server.Router.POST("/users/:userID/friend-requests", friendRequestController.SendFriendRequest)

		server.Router.POST("/users/:userID/status", userController.GetUserStatus)

		// 获取用户的所有好友请求
		server.Router.GET("/users/:userID/friend-requests", friendRequestController.GetFriendApplicationList)

		// 同意好友请求
		server.Router.PUT("/users/:userID/friend-requests/:requestID/accept", friendRequestController.AcceptFriendRequest)

		// 拒绝好友请求
		server.Router.PUT("/users/:userID/friend-requests/:requestID/reject", friendRequestController.RejectFriendRequest)

		// 删除好友请求
		server.Router.DELETE("/users/:userID/friend-requests/:requestID", friendRequestController.CancelFriendRequest)

		// 获取用户的所有好友
		server.Router.GET("/users/:userID/friends", friendshipController.GetFriends)

		// 添加好友
		server.Router.POST("/users/:userID/friends/:friendID", friendshipController.AddFriend)

		// 删除好友
		server.Router.DELETE("/users/:userID/friends/:friendID", friendshipController.RemoveFriend)

		// 更新好友备注
		server.Router.PUT("/users/:userID/friends/:friendID/remark", friendshipController.UpdateFriendRemark)

		// 获取好友备注
		server.Router.GET("/users/:userID/friends/:friendID/remark", friendshipController.GetFriendRemark)

		//TODO
		server.Router.POST("/users/friendsInfo", friendshipController.GetFriendsInfo)
		server.Router.POST("/users/:userID/checkFriend", friendshipController.CheckFriend)

		server.Router.POST("/users/:userID/joinGroupList", GroupController.GetJoinedGroupList)

		// 启动HTTP服务
		go server.Start()

		// 启动Socket服务
		go socket.Start()
	})
	if err != nil {
		log.Fatal(err)
	}
}
