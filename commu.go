package iotbase

type ICommunicator interface {
	Send(cmd InstrCommand) error
}
