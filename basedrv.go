package iotbase

import "encoding/json"

type BaseDriver struct {
	device        Device
	model         ModelAriot
	deploy        DeployAriot
	handler       IEventHandler
	initParamHook IParamParseHook
}

func (drv *BaseDriver) Init(device Device, configJson string, modelJson string, handler IEventHandler) error {
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

func (drv *BaseDriver) emitStatus(status Status) {
	if drv.handler != nil {
		drv.handler.onStatus(status)
	}
}

func (drv *BaseDriver) emitError(warn Warn) {
	if drv.handler != nil {
		drv.handler.onError(warn)
	}
}

func (drv *BaseDriver) emitProperty(property Property) {
	if drv.handler != nil {
		drv.handler.onProperty(property)
	}
}

func (drv *BaseDriver) emitProperties(properties []Property) {
	if drv.handler != nil {
		drv.handler.onProperties(properties)
	}
}

func (drv *BaseDriver) emitCommandReply(reply CommandReply) {
	if drv.handler != nil {
		drv.handler.onCommandReply(reply)
	}
}
