package iotbase

import "github.com/akingbrDu/iotbase/core"

type DbgInformation struct {
	FileName   string
	MethodName string
	Line       int
	Reason     string
}

type IEventHandler interface {
	onError(warn core.Warn)
	onProperty(property core.Property)
	onProperties(properties []core.Property)
	onStatus(status core.Status)
	onCommandReply(reply core.CommandReply)
}
