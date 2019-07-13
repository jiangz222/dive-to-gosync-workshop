package sync

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

// 并发的 对某一个特定的 f func() 调用Do时，利用原子操作和mutex，保证 只调用一次callback
func (o *Once) Do(f func() error) error {
	// 这里不是必须吧？
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}

	return o.slowDo(f)
}

func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
