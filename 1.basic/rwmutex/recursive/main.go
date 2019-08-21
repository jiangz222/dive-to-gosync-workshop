package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(1)

	// 后面的递归过程中制造一个Lock
	go func() {
		defer wg.Done()

		time.Sleep(200 * time.Millisecond)
		fmt.Println("try Lock")
		// lock 被递归的前2个rlock阻塞，第3个rlock被本lock阻塞，前2个rlock需要第三个rlock的unlock才能解锁, 构成死锁
		mu.Lock()
		fmt.Println("Lock")
		time.Sleep(100 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Unlock")
	}()

	go func() {
		defer wg.Done()
		factorial(&mu, 3)
	}()
	wg.Wait()
}

func factorial(m *sync.RWMutex, n int) int {
	if n < 1 {
		return 0
	}
	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n
}
