package main

import (
	"fmt"
	"log"
	"time"
)

// [jz] 关于timer一个比较重要的点是，newtimer后，timer会放入最小堆，然后有一个goroutine来扫描，到期的进行回调和channel写入
// 	stop只负责将timer从堆删除，不负责close channel
func main() {
	log.Println("✔︎ resetBeforeFired")
	resetBeforeFired()
	fmt.Println()

	log.Println("✘ wrongResetAfterFired")
	wrongResetAfterFired()
	fmt.Println()

	log.Println("✔︎ correctResetAfterFired")
	correctResetAfterFired()
	fmt.Println()

	log.Println("✔︎ stop n times")
	stopMore()
	fmt.Println()

	log.Println("✘ stop n times but with drain")
	wrongStopMore()
	fmt.Println()

	log.Println("✘ too many receiving")
	wrongReceiveMore()
}

func resetBeforeFired() {
	timer := time.NewTimer(5 * time.Second)
	b := timer.Stop()
	log.Printf("stop: %t", b)
	timer.Reset(1 * time.Second)
	t := <-timer.C
	log.Printf("fired at %s", t.String())
}

func wrongResetAfterFired() {
	timer := time.NewTimer(5 * time.Millisecond)
	time.Sleep(time.Second) // sleep 1s能保证上面的timer 超时，channel被写入

	b := timer.Stop()
	log.Printf("stop: %t", b)
	tt := timer.Reset(10 * time.Second)
	fmt.Println(tt)
	// 此时拿到的是第一个timer（5毫秒那个）的timeout的channel值
	t := <-timer.C
	log.Printf("fired at %s", t.String())
}

func correctResetAfterFired() {
	timer := time.NewTimer(5 * time.Millisecond)
	time.Sleep(time.Second)

	b := timer.Stop()
	log.Printf("stop: %t", b)
	// 如果stop的时候发现已经超时，此时要把channel里的写入读出，免得后面reset时读出之前的channel里的值
	if !b {
		t := <-timer.C
		fmt.Println(t.String())
	}
	log.Printf("reset")
	timer.Reset(10 * time.Second)
	t := <-timer.C
	log.Printf("fired at %s", t.String())
}

func wrongReceiveMore() {
	timer := time.NewTimer(5 * time.Millisecond)
	t := <-timer.C
	log.Printf("fired at %s", t.String())

	t = <-timer.C
	log.Printf("receive again at %s", t.String())
}

func stopMore() {
	timer := time.NewTimer(5 * time.Millisecond)
	b := timer.Stop()
	log.Printf("stop: %t", b)
	time.Sleep(time.Second)
	b = timer.Stop()
	log.Printf("stop more: %t", b)
}

/*
	newtimer后，timer会放入最小堆，然后有一个goroutine来扫描，到期的进行回调和channel写入
 	stop只负责将timer从堆删除，不负责close channel
*/
func wrongStopMore() {
	timer := time.NewTimer(5 * time.Millisecond)
	b := timer.Stop()
	log.Printf("stop: %t", b)
	time.Sleep(time.Second)
	b = timer.Stop()
	if !b { // 可以考虑这样解决：if !b && len(timer.C) > 0
		// 之所以出问题，是因为，第一次Stop调用，发生在timer超时前，此时timer已经从堆删除，而timer本身没有超时，所以不需要发送channel
		// 此时你去等待timer.C是不会有结果的
		// 比如你在第一个timer.Stop前sleep 1s，让timer超时，channel会被写入，此时等待timer	.C就不会有问题
		<-timer.C
	}
	time.Sleep(1 * time.Second)
	log.Printf("stop more: %t", b)
}
