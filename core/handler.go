package core

type DbgInformation struct {
	FileName   string
	MethodName string
	Line       int
	Reason     string
}

type IEventHandler interface {
	OnError(warn Warn)
	OnProperty(property Property)
	OnProperties(properties []Property)
	OnStatus(status Status)
	OnCommandReply(reply CommandReply)
}
