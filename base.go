package iotbase

import (
	"math/rand"
	"time"
)

type ICommunication interface {
	Start() error
}

type ProtocolBase struct {
	ProtocolName    string
	ProtocolVersion string
	ProtocolType    int
}

func GetRandom() int {
	rand.Seed(time.Now().Unix())
	rn := 10
	return rand.Intn(rn)
}
