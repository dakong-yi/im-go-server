package message

type ReqHandler interface {
	Handle(req MessageRequest)
}
