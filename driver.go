package session

import (
	"time"

	. "github.com/chefsgo/base"
)

type (
	// Driver 数据驱动
	Driver interface {
		Connect(name string, config Config) (Connect, error)
	}

	// Connect 会话连接
	Connect interface {
		Open() error
		Close() error

		Read(id string) (Map, error)
		Write(id string, value Map, expiry time.Duration) error
		Delete(id string) error
		Clear(perfix string) error
	}
)

// Driver 注册驱动
func (module *Module) Driver(name string, driver Driver, override bool) {
	module.mutex.Lock()
	defer module.mutex.Unlock()

	if driver == nil {
		panic("Invalid session driver: " + name)
	}

	if override {
		module.drivers[name] = driver
	} else {
		if module.drivers[name] == nil {
			module.drivers[name] = driver
		}
	}
}
