package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan struct{}, 1)

	select {
	// Q：两个case和default怎么走？？？
	case ch <- struct{}{}:
		fmt.Println("1")
	case ch <- struct{}{}:
		fmt.Println("2")
	default:
		fmt.Println("hihi")
	}
	time.Sleep(1 * time.Second)

}
