package sync

import "sync/atomic"

type Flag struct {
	f uint32
}

func (a *Flag) Set() {
	atomic.StoreUint32(&a.f, 1)
}

func (a *Flag) Clear() {
	atomic.StoreUint32(&a.f, 0)
}

func (a *Flag) IsSet() bool {
	return atomic.LoadUint32(&a.f) == 1
}

func (a *Flag) String() string {
	if a.IsSet() {
		return "true"
	}
	return "false"
}
