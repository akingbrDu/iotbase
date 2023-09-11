package iotbase

import "encoding/json"

type baseDriver struct {
	Model ModelAriot
}

func (drv *baseDriver) ParseModel(modelJson string) error {
	err := json.Unmarshal([]byte(modelJson), &drv.Model)
	if err != nil {
		return err
	}
	return nil
}

func (drv *baseDriver) ParseDeploy(configJson string) error {
	return nil
}

func (drv *baseDriver) Init(configJson string, modelJson string) error {
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
