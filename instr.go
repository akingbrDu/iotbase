package iotbase

type IInstrResultParser interface {
	parseContent(interface{}) map[string]interface{}
}

type InstrCommand struct {
	parser IInstrResultParser
}

type BinaryInstrCommand struct {
	InstrCommand
	data []byte
}

type StringInstrCommand struct {
	InstrCommand
	data string
}
