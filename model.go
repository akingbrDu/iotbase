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

/* 物模型结构体定义最早Json都是首字母大写
后来的版本新增了一些字段，有很多小写开头
不太统一 -Syg*/

type EnumDataSpec struct {
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type DataTypeSpec struct {
	Min      string         `json:"Min"`
	Max      string         `json:"Max"`
	Step     string         `json:"Step"`
	UnitName string         `json:"UnitName"`
	Length   int            `json:"length"`
	Bool0    string         `json:"bool0"`
	Bool1    string         `json:"bool1"`
	EnumList []EnumDataSpec `json:"enumList"`
}

type ModelDataType struct {
	Type string       `json:"Type"`
	Spec DataTypeSpec `json:"Spec"`
}

type ModelParam struct {
	Identifier string        `json:"Identifier"`
	Name       string        `json:"Name"`
	DataType   ModelDataType `json:"DataType"`
}

type ModelProperty struct {
	Id          string        `json:"Id"`
	Identifier  string        `json:"Identifier"`
	Name        string        `json:"Name"`
	AccessMode  int           `json:"AccessMode"`
	Description string        `json:"Description"`
	DataType    ModelDataType `json:"DataType"`
}

type ModelEvent struct {
	Id           string       `json:"Id"`
	Identifier   string       `json:"Identifier"`
	Name         string       `json:"Name"`
	EventType    string       `json:"EventType"`
	OutputParams []ModelParam `json:"OutputParams"`
}

type ModelCommand struct {
	Id           string       `json:"Id"`
	Identifier   string       `json:"Identifier"`
	Name         string       `json:"Name"`
	CommandType  string       `json:"commandType"`
	Description  string       `json:"Description"`
	InputParams  []ModelParam `json:"InputParams"`
	OutputParams []ModelParam `json:"OutputParams"`
}

type ModelAriot struct {
	Properties []ModelProperty `json:"Property"`
	Events     []ModelEvent    `json:"Event"`
	Commands   []ModelCommand  `json:"Command"`
}
