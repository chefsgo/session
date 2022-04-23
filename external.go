package session

import (
	"time"

	. "github.com/chefsgo/base"
)

func Read(id string) (Map, error) {
	return module.Read(id)

}

func Write(id string, value Map, expiries ...time.Duration) error {
	return module.Write(id, value, expiries...)
}

func Delete(id string) error {
	return module.Delete(id)
}

func Clear() error {
	return module.Clear()
}
