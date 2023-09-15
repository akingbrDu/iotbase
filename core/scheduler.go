package core

type IScheduler interface {
	Register(query IQuery) error
	Run() error
}

// IQuery 查询类驱动需要支持
type IQuery interface {
	Query() error
	GetFreq() int
}

// ICallback 回调类驱动需要支持
type ICallback interface {
	Callback() error
}

// DefaultSharedScheduler 默认共享端口查询调度
type DefaultSharedScheduler struct {
}

func (s *DefaultSharedScheduler) Register(query IQuery) error {
	return nil
}

func (s *DefaultSharedScheduler) Run() error {
	return nil
}
