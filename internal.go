package session

import (
	"time"

	. "github.com/chefsgo/base"
)

func (this *Module) Read(id string) (Map, error) {
	locate := this.hashring.Locate(id)

	if inst, ok := this.instances[locate]; ok {
		key := inst.config.Prefix + id //加前缀
		return inst.connect.Read(key)
	}

	return nil, errInvalidSessionConnection

}

func (this *Module) Write(id string, value Map, expiries ...time.Duration) error {
	locate := this.hashring.Locate(id)

	if inst, ok := this.instances[locate]; ok {
		expiry := inst.config.Expiry
		if len(expiries) > 0 {
			expiry = expiries[0]
		}

		//KEY加上前缀
		key := inst.config.Prefix + id

		return inst.connect.Write(key, value, expiry)
	}

	return errInvalidSessionConnection
}

func (this *Module) Delete(id string) error {
	locate := this.hashring.Locate(id)

	if inst, ok := this.instances[locate]; ok {
		key := inst.config.Prefix + id
		return inst.connect.Delete(key)
	}

	return errInvalidSessionConnection
}

func (this *Module) Clear() error {
	for _, inst := range this.instances {
		inst.connect.Clear(inst.config.Prefix)
	}

	return errInvalidSessionConnection
}
