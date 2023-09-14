package iotbase

import (
	"encoding/json"
	"github.com/akingbrDu/iotbase/core"
)

type BaseDriver struct {
	device        core.Device
	model         core.ModelAriot
	deploy        core.DeployAriot
	handler       IEventHandler
	initParamHook IParamParseHook
}

func (drv *BaseDriver) Init(device core.Device, configJson string, modelJson string, handler IEventHandler) error {
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
		drv.handler.onStatus(status)
	}
}

func (drv *BaseDriver) emitError(warn core.Warn) {
	if drv.handler != nil {
		drv.handler.onError(warn)
	}
}

func (drv *BaseDriver) emitProperty(property core.Property) {
	if drv.handler != nil {
		drv.handler.onProperty(property)
	}
}

func (drv *BaseDriver) emitProperties(properties []core.Property) {
	if drv.handler != nil {
		drv.handler.onProperties(properties)
	}
}

func (drv *BaseDriver) emitCommandReply(reply core.CommandReply) {
	if drv.handler != nil {
		drv.handler.onCommandReply(reply)
	}
}
