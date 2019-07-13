package main

import (
	"time"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{} // 先往ch发送，保证下一个lock的可以不阻塞
	return mu
}
func (m *Mutex) Lock() {
	<-m.ch
}
func (m *Mutex) Unlock() {
	select {
	// 先往ch发送，保证下一个lock的可以不阻塞
	// 如果没有lock，则说明ch里已经有一个消息，不能再往ch里面塞消息，即写不成功，所以这里会走到default
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}
func (m *Mutex) TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-time.After(timeout):
	}
	return false
}
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}
func (m *Mutex) Len() int {
	return len(m.ch)
}
