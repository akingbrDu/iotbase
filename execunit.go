package iotbase

import "github.com/akingbrDu/iotbase/core"

// 驱动执行单元
// 执行单元调度类型
const (
	Query = iota
	Callback
)

// 执行单元执行模式
const (
	Interval = iota
	Once
)

type ExecUnit struct {
	RunType int
	RunMode int
	Freq    int
	Sender  ICommunicator
	CmdSet  []core.InstrCommand
}
