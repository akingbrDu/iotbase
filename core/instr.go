package core

//type IInstrResultParser interface {
//	ParseContent(interface{}) map[string]interface{}
//}

type IInstrCommand interface {
	CheckIfFinished() bool
	Parse() map[string]interface{}
}

type BaseInstrCommand struct {
	//Parser IInstrResultParser
}

type BinaryInstrCommand struct {
	BaseInstrCommand
	OutData      []byte
	IncomingData []byte
}

func (cmd *BinaryInstrCommand) checkIfFinished() bool {
	return true
}

type StringInstrCommand struct {
	BaseInstrCommand
	OutData      string
	IncomingData string
}

func (cmd *StringInstrCommand) checkIfFinished() bool {
	return true
}
