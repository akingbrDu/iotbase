package iotbase

type DbgInformation struct {
	FileName   string
	MethodName string
	Line       int
	Reason     string
}

type IEventHandler interface {
	onError(warn Warn)
	onProperty(property Property)
	onProperties(properties []Property)
	onStatus(status Status)
	onCommandReply(reply CommandReply)
}
