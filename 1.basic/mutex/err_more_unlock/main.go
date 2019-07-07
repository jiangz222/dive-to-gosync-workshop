package main

import (
	"sync"
	"time"
)
/**
	mu 的lock与unlock即对mu.state进行原子加一减一
 */
func main() {
	var mu sync.Mutex
	go func() {
		mu.Lock()
		time.Sleep(5 * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)
	//fatal error: sync: unlock of unlocked mutex
	mu.Unlock()

	select {}
}
