package iotbase

import "math/rand"

type ICommunication interface {
	Start() error
}

type ProtocolBase struct {
	ProtocolName    string
	ProtocolVersion string
	ProtocolType    int
}

func GetRandom() int {
	rn := 10
	return rand.Intn(rn)
}
