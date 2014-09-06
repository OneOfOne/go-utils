package spinlock

import (
	"runtime"
	"sync/atomic"
)

var state = [2]string{"Unlocked", "Locked"}

// SpinLock implements a simple atomic spin lock, the zero value for a SpinLock is an unlocked spinlock.
type SpinLock struct {
	l uint32
}

// Lock locks sl. If the lock is already in use, the caller blocks until Unlock is called
func (sl *SpinLock) Lock() {
	for !atomic.CompareAndSwapUint32(&sl.l, 0, 1) {
		runtime.Gosched() //allow other goroutines to do stuff.
	}
}

// Unlock unlocks sl, unlike [Mutex.Unlock](http://golang.org/pkg/sync/#Mutex.Unlock),
// there's no harm calling it on an unlocked SpinLock
func (sl *SpinLock) Unlock() {
	atomic.StoreUint32(&sl.l, 0)
}

// TryLock will try to lock sl and return whether it succeed or not without blocking.
func (sl *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&sl.l, 0, 1)
}

func (sl *SpinLock) String() string {
	return state[sl.l]
}
