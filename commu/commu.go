package commu

import _ "github.com/akingbrDu/iotbase/core"

type ICommunicator interface {
	Send(cmd InstrCommand) error
}
