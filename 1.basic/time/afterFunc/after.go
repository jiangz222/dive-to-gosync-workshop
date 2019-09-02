package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second, func() {
		// [jz]timer.C用来触发本func()，阻塞于此直到timeout
		fmt.Println("fired")
	})
	t := <-timer.C                        // nil
	log.Printf("fired at %s", t.String()) // 海枯石烂你也等不来
}
