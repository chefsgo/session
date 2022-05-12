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
