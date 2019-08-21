package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type RWMutex struct {
	sync.RWMutex
}

type m struct {
	w           sync.Mutex
	writerSem   uint32
	readerSem   uint32
	readerCount int32
	readerWait  int32
}

func (rw *RWMutex) ReaderCount() int {
	// [jz] 可以这样访问到结构小写的成员...
	v := (*m)(unsafe.Pointer(&rw.RWMutex))
	c := int(v.readerCount)
	if c < 0 {
		c = int(v.readerWait)
	}

	return c
}

func (rw *RWMutex) WriterCount() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&rw.RWMutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return int(v)
}

func main() {
	var mu RWMutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.RLock()
			fmt.Println("rlock")
			time.Sleep(time.Hour)
			mu.RUnlock()
		}()
	}

	for i := 0; i < 50; i++ {
		go func() {
			mu.Lock()
			fmt.Println("write lock")
			time.Sleep(time.Hour)
			mu.Unlock()
		}()
	}

	fmt.Println("readers: ", mu.ReaderCount())
	// write lock required，so the count is increased
	fmt.Println("writer: ", mu.WriterCount())
	time.Sleep(10 * time.Second)
}
