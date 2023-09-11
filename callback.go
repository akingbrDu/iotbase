package iotbase

type DbgInformation struct {
	FileName   string
	MethodName string
	Line       int
	Reason     string
}

type IEventHandler interface {
	onError(msg string, detail DbgInformation)
}
