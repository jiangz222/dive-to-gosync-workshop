package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("fired")
	})
	fmt.Println("after timer", timer.C)
	// [jz]AfterFunc 生成的timer不会调用C，没有初始化C
	t := <-timer.C                        // nil
	log.Printf("fired at %s", t.String()) // 海枯石烂你也等不来
}
