package wrong

import "sync/atomic"

type Once struct {
	done uint32
}

// 怎么wrong了？

func (o *Once) Do(f func()) {
	if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		return
	}
	f()
}