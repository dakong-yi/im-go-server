package request

import (
	"github.com/dakong-yi/im-go-server/pkg/message"
	"github.com/dakong-yi/im-go-server/pkg/user"
)

// 定义处理请求的接口
type requestHandler interface {
	handle(data string, user *user.User) message.MessageRequest
}
