package iotbase

const (
	INT32  = "int32"
	UINT32 = "int32"
	FLOAT  = "float"
	DOUBLE = "double"
	STRING = "string"
	BOOL   = "bool"
	ENUM   = "enum"
	JSON   = "json"
)

const (
	Info    = 0x01
	Warning = 0x02
	Fault   = 0x04
	Unknown = 0x08
)

const (
	sync  = 0x01
	async = 0x02
)

func ToWarningType(typeValue int) string {
	switch typeValue {
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Fault:
		return "Fault"
	default:
		return "Unknown"
	}
}

func ToSyncType(typeValue int) string {
	switch typeValue {
	case sync:
		return "sync"
	case async:
		return "async"
	default:
		return "Unknown"
	}
}

type EnumDataSpec struct {
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type DataTypeSpec struct {
	Min      int            `json:"min"`
	Max      int            `json:"max"`
	Step     string         `json:"step"`
	UnitName string         `json:"unitName"`
	Length   int            `json:"length"`
	Bool0    string         `json:"bool0"`
	Bool1    string         `json:"bool1"`
	EnumList []EnumDataSpec `json:"enumList"`
}

type ModelDataType struct {
	Type string       `json:"type"`
	Spec DataTypeSpec `json:"spec"`
}

type ModelParam struct {
	Identifier string        `json:"identifier"`
	Name       string        `json:"name"`
	DataType   ModelDataType `json:"dataType"`
}

type ModelProperty struct {
	Id          string        `json:"id"`
	Identifier  string        `json:"identifier"`
	Name        string        `json:"name"`
	AccessMode  string        `json:"accessMode"`
	Description string        `json:"description"`
	DataType    ModelDataType `json:"dataType"`
}

type ModelEvent struct {
	Id           string       `json:"id"`
	Identifier   string       `json:"identifier"`
	Name         string       `json:"name"`
	EventType    string       `json:"eventType"`
	OutputParams []ModelParam `json:"outputParams"`
}

type ModelCommand struct {
	Id           string       `json:"id"`
	Identifier   string       `json:"identifier"`
	Name         string       `json:"name"`
	CommandType  string       `json:"commandType"`
	Description  string       `json:"description"`
	InputParams  []ModelParam `json:"inputParams"`
	OutputParams []ModelParam `json:"outputParams"`
}

type ModelAriot struct {
	Properties []ModelProperty
	Events     []ModelEvent
	Commands   []ModelCommand
}
