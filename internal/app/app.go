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
	container.Provide(service.NewMessageService)
	container.Provide(service.NewConversationService)
	container.Provide(repository.NewMessageRepoImpl)
	container.Provide(repository.NewConversationRepoImpl)
	container.Provide(repository.NewUserConversationRepoImpl)
	// 增加其他依赖...

	err := container.Invoke(func(server *http.Server, socket *socket.SocketServer,
		conversationController *controller.ConversationController,
		messageController *controller.MessageController,
		// 增加其他依赖...
	) {
		// 在这里可以使用注入后的控制器注册路由
		server.Router.POST("/conversation", conversationController.CreateConversation)
		server.Router.GET("/conversation/:conversationID", conversationController.GetConversation)
		server.Router.GET("/conversation/users/:userID", conversationController.GetConversationsByUserID)

		// 注册路由和处理函数
		server.Router.POST("/message", messageController.CreateMessage)
		server.Router.GET("/message/:conversationID", messageController.GetMessagesByConversationID)
		// 启动HTTP服务
		// go server.Start()

		// 启动Socket服务
		go socket.Start()
	})
	if err != nil {
		log.Fatal(err)
	}
}
