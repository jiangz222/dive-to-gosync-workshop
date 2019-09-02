package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// [jz] ticker是一直在自动跑，循环定时器，一次到时候下一次马上开始运行
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	log.Println("create a ticker")

	t := <-ticker.C
	log.Println("the first tick:", t.String())

	time.Sleep(2 * time.Second)
	t = <-ticker.C // chan吐出来的时间是生效时间不是当前时间，说明ticker是马上生效的
	log.Println("then the second tick:", t.String(), time.Now())

	t = <-ticker.C
	log.Println("then the third tick:", t.String(), time.Now())

	//[jz] 循环定时器更好的例子：
	tt := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tt.C:
			fmt.Println("ticker run :", time.Now())
			time.Sleep(3 * time.Second)
			// sleep了3秒后，下一次ticker是2秒生效，说明ticker是一个循环定时器
		}
	}

}
