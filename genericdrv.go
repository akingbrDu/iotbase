package iotbase

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GenericDriver struct {
	BaseDriver
	freq int
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
// 都统一放到base包的最下层
func (drv *GenericDriver) Init(device Device, configJson string, modelJson string, handler IEventHandler) error {
	if drv.initParamHook != nil {
		if err := drv.initParamHook.onParamBefore(); err != nil {
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

	commuParamInfo, err := drv.parseRawMessage(&drv.deploy.Communication.Param)
	if err != nil {
		fmt.Println("fail to decode communication param field in json with err: :", err)
		return err
	}

	freq, ok := runtimeInfo["freq"].(int)
	if !ok {
		fmt.Println("can't parse freq field in json\"")
		return errors.New("can't parse freq field in json")
	}
	drv.freq = freq

	if drv.initParamHook != nil {
		if err := drv.initParamHook.onParamParsed(); err != nil {
			return err
		}
	}

	defCommunicator, err := drv.createCommunicator(drv.deploy.Communication.Type, commuParamInfo)
	if err != nil {
		fmt.Println("fail to create communicator with err: :", err)
		return err
	}

	//创建驱动执行单元
	units, err := drv.createExecUnits(defCommunicator)
	if err != nil {
		return err
	}

	//分配调度
	err = drv.schedule(units)
	if err != nil {
		return err
	}

	return nil
}

func (drv *GenericDriver) createCommunicator(cType string, param map[string]interface{}) (ICommunicator, error) {
	if drv.initParamHook != nil {
		if communicator, err := drv.initParamHook.onCreateDefaultCommunicator(); err != nil {
			return nil, err
		} else if communicator != nil {
			return communicator, nil
		}
	} else {
		//create default	waiting..
	}
	return nil, nil
}

func (drv *GenericDriver) createExecUnits(defCommunicator ICommunicator) ([]ExecUnit, error) {
	var units []ExecUnit
	if drv.initParamHook != nil {
		if err := drv.initParamHook.onCreateExecUnits(units); err != nil {
			return nil, err
		}

		//如果不自定义通讯器，则使用默认通讯器
		for _, unit := range units {
			if unit.Sender == nil {
				unit.Sender = defCommunicator
			}

			if unit.RunType == Query && unit.Freq <= 0 {
				unit.Freq = drv.freq
			}
		}
	}
	return units, nil
}

func (drv *GenericDriver) schedule(units []ExecUnit) error {
	return nil
}
