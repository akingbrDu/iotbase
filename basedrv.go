package iotbase

import (
	"encoding/json"
	"github.com/akingbrDu/iotbase/core"
)

type BaseDriver struct {
	device        core.Device
	model         core.ModelAriot
	deploy        core.DeployAriot
	handler       core.IEventHandler
	initParamHook core.IInitHook
}

func (drv *BaseDriver) Init(device core.Device, configJson string, modelJson string, handler core.IEventHandler) error {
	drv.handler = handler
	drv.device = device

	err := json.Unmarshal([]byte(modelJson), &drv.model)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(configJson), &drv.deploy)
	if err != nil {
		return err
	}

	return nil
}

func (drv *BaseDriver) emitStatus(status core.Status) {
	if drv.handler != nil {
		drv.handler.OnStatus(status)
	}
}

func (drv *BaseDriver) emitError(warn core.Warn) {
	if drv.handler != nil {
		drv.handler.OnError(warn)
	}
}

func (drv *BaseDriver) emitProperty(property core.Property) {
	if drv.handler != nil {
		drv.handler.OnProperty(property)
	}
}

func (drv *BaseDriver) emitProperties(properties []core.Property) {
	if drv.handler != nil {
		drv.handler.OnProperties(properties)
	}
}

func (drv *BaseDriver) emitCommandReply(reply core.CommandReply) {
	if drv.handler != nil {
		drv.handler.OnCommandReply(reply)
	}
}
