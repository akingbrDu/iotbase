package iotbase

import "encoding/json"

type BaseDriver struct {
	Model ModelAriot
}

func (drv *BaseDriver) ParseModel(modelJson string) error {
	err := json.Unmarshal([]byte(modelJson), &drv.Model)
	if err != nil {
		return err
	}
	return nil
}

func (drv *BaseDriver) ParseDeploy(configJson string) error {
	return nil
}

func (drv *BaseDriver) Init(configJson string, modelJson string) error {
	err := drv.ParseModel(modelJson)
	if err != nil {
		return err
	}

	err = drv.ParseDeploy(configJson)
	if err != nil {
		return err
	}
	return nil
}
