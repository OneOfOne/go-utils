package sync

import "sync/atomic"

type AtomicFlag struct {
	v uint32
}

func (a *AtomicFlag) Set() bool {
	return atomic.CompareAndSwapUint32(&a.v, 0, 1)
}

func (a *AtomicFlag) Clear() bool {
	return atomic.CompareAndSwapUint32(&a.v, 1, 0)
}

func (a *AtomicFlag) IsSet() bool {
	return atomic.LoadUint32(&a.v) == 1
}

func (a *AtomicFlag) String() string {
	if a.IsSet() {
		return "true"
	}
	return "false"
}
