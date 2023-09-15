package core

const (
	ModBusRTU   string = "ModbusRTU"
	ModBusTCP   string = "ModbusTCP"
	ModBusAscii string = "ModbusAscii"
	BacnetIp    string = "BACnet"
	Dlt645_07   string = "DL/T645-2007"
)

// 基本结构体 below

type Device struct {
	OwnerSn string `json:"ownerSn"`
	DevSn   string `json:"devSn"`
}

type CommandParam struct {
	Identifier string      `json:"identifier"`
	Value      interface{} `json:"value"`
}

type Command struct {
	Identifier string `json:"identifier"`
	DevSn      string `json:"devSn"`
	Params     []CommandParam
}

type CommandReply struct {
	Identifier string         `json:"identifier"`
	Code       int            `json:"code"`
	DevSn      string         `json:"devSn"`
	Params     []CommandParam `json:"params"`
}

type Property struct {
	ItemSet map[string]interface{}
}

type Status struct {
	State      int         `json:"state"`
	Identifier string      `json:"identifier"`
	Value      interface{} `json:"value"`
}

type WarningParam struct {
	ParamIdentifier string      `json:"paramIdentifier"`
	ParamValue      interface{} `json:"paramValue"`
}

type Warn struct {
	Classify    int    `json:"classify"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}

type DeviceWarn struct {
	Warn
	DevSn      string         `json:"devSn"`
	Identifier string         `json:"identifier"`
	ParamList  []WarningParam `json:"paramList"`
}

// IErrorHandler 定义处理器
type IErrorHandler interface {
}
