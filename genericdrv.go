package iotbase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/akingbrDu/iotbase/core"
)

type GenericDriver struct {
	BaseDriver
	freq               int
	runtimeInfo        map[string]interface{}
	cmuType            string
	cmuParamInfo       map[string]interface{}
	protoType          string
	protoParamInfo     map[string]interface{}
	defSharedScheduler core.IScheduler
}

func (drv *GenericDriver) parseRawMessage(rawMessage *json.RawMessage) (map[string]interface{}, error) {
	var msgJson interface{}
	err := json.Unmarshal(*rawMessage, &msgJson)
	if err != nil {
		return nil, err
	}

	msgInfo, ok := msgJson.(map[string]interface{})
	if !ok {
		return nil, errors.New("can't parse raw message field in json")
	}

	return msgInfo, nil
}

// Init  因为Go支持组合，对abstract继承的支持不是很好，所以所有Hook接口
// 都统一放到Generic实现的最下面
func (drv *GenericDriver) Init(device core.Device, configJson string, modelJson string, handler core.IEventHandler) error {
	if drv.initParamHook != nil {
		if err := drv.initParamHook.OnBefInitialize(); err != nil {
			return err
		}
	}

	err := drv.BaseDriver.Init(device, configJson, modelJson, handler)
	if err != nil {
		return err
	}

	runtimeInfo, err := drv.parseRawMessage(&drv.deploy.Runtime)
	if err != nil {
		fmt.Println("fail to decode runtime field in json with err: :", err)
		return err
	}

	cmuParamInfo, err := drv.parseRawMessage(&drv.deploy.Communication.Param)
	if err != nil {
		fmt.Println("fail to decode communication param field in json with err: :", err)
		return err
	}

	protoParamInfo, err := drv.parseRawMessage(&drv.deploy.Proto.Param)
	if err != nil {
		fmt.Println("fail to decode proto param field in json with err: :", err)
		return err
	}

	freq, ok := runtimeInfo["freq"].(int)
	if !ok {
		fmt.Println("can't parse freq field in json\"")
		return errors.New("can't parse freq field in json")
	}

	//解析一些常用的字段
	drv.runtimeInfo = runtimeInfo
	drv.cmuParamInfo = cmuParamInfo
	drv.cmuType = drv.deploy.Communication.Type
	drv.protoParamInfo = protoParamInfo
	drv.protoType = drv.deploy.Proto.Type
	drv.freq = freq

	drv.CreateScheduler()

	if drv.initParamHook != nil {
		if err := drv.initParamHook.OnInitialized(); err != nil {
			return err
		}
	}

	return nil
}

func (drv *GenericDriver) CreateScheduler() {
	drv.defSharedScheduler = &core.DefaultSharedScheduler{}
}

func (drv *GenericDriver) Run() error {
	err := drv.defSharedScheduler.Run()
	if err != nil {
		return err
	}
	return nil
}
