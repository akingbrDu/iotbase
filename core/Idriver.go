package core

type IDriver interface {
	// Init 	 *
	//@brief: init
	//	 * @param: configJson: 配置参数
	//	 * @param: modelJson: 物模型参数
	//	 * @return:   错误码
	Init(device Device, configJson string, modelJson string, handler IEventHandler) error

	// Run  *
	// 驱动启动运行
	Run() error

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

type IQueryDriver interface {
	IDriver
	IQuery
}

type ICallbackDriver interface {
	IDriver
	ICallback
}

type IQueryAndCallbackDriver interface {
	IDriver
	IQuery
	ICallback
}

// IInitHook  初始化解析钩子
type IInitHook interface {
	// OnBefInitialize  *初始化解析参数前的钩子
	OnBefInitialize() error
	// OnInitialized  *初始化解析参数后的钩子
	OnInitialized() error
}
