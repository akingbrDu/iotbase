package iotbase

type IDriver interface {
	// Init 	 *
	//@brief: init
	//	 * @param: configJson: 配置参数
	//	 * @param: modelJson: 物模型参数
	//	 * @return:   错误码
	Init(device Device, configJson string, modelJson string, handler IEventHandler) error

	//// Query  *
	////@brief: start
	////	 * @param:
	////	 * @return:   字节流 错误码
	//Query() ([]byte, error)
	//
	//// Process  *
	////@brief: process
	////	 * @param: data: 字节流
	////	 * @return:   上报数据, 错误码
	//Process(data []byte) (string, error)

	// ExeCommand  *
	//@brief: execute command
	//	 * @param: cmd: 命令结构
	//	 * @return:   错误码
	ExeCommand(cmd Command) error

	// SetOption  *
	//@brief: set option
	//	 * @param: key: 键
	//	 * @param: value: 值
	//	 * @return:   错误码
	SetOption(key string, value string) error
}

// IParamParseHook  初始化解析钩子
type IParamParseHook interface {
	// onParamBefore  *初始化解析参数前的钩子
	onParamBefore() error
	// onParamParsed  *初始化解析参数后的钩子
	onParamParsed() error
	// onCreateDefaultCommunicator 创建默认通讯器
	onCreateDefaultCommunicator() (ICommunicator, error)
	// onCreateExecUnits 创建执行单元
	onCreateExecUnits(units []ExecUnit) error
}
