package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"time"
)

// parent的cancel可以取消子context.
// 子context只能取消自己和它的子.
// 子会propagateCancel，一直上上找，直到找到一个可cancel的context, 把自己加到它的children中.
func main() {
	ctx := context.Background()

	ctx1, c1 := context.WithCancel(ctx)
	go func() {
		fmt.Println("g1 start")
		<-ctx1.Done()
		fmt.Println("g1 done, err:", ctx.Err())
	}()

	ctx2, c2 := context.WithCancel(ctx1)
	go func() {
		fmt.Println("g2 start")
		<-ctx2.Done()
		fmt.Println("g2 done, err:", ctx1.Err())
	}()

	ctx3, c3 := context.WithCancel(ctx2)
	go func() {
		fmt.Println("g3 start")
		<-ctx3.Done()
		fmt.Println("g3 done, err:", ctx2.Err())
	}()

	//time.Sleep(1e9)
	//fmt.Println("try to  call c1")
	//c1()
	//fmt.Println("c1 called")
	//time.Sleep(5 * time.Second)
	//fmt.Println("try to  call c2")
	//c2()
	//fmt.Println("c2 called")
	//time.Sleep(5 * time.Second)
	//fmt.Println("try to  call c3")
	//c3()
	//fmt.Println("c3 called")
	/*
		   output:
			g1 start
		g3 start
		g2 start
		   try to  call c1
		   c1 called
		   g1 done, err: <nil>
			// g2和g3的done次序是比较随机的
		   g2 done, err: context canceled
		   g3 done, err: context canceled
		 	// c1 取消后，其子context都取消了
		   try to  call c2
		   c2 called
		   try to  call c3
		   c3 called
	*/
	//
	time.Sleep(1e9)
	fmt.Println("try to  call c3")
	c3()
	fmt.Println("c3 called")
	time.Sleep(5 * time.Second)
	fmt.Println("try to  call c2")
	c2()
	fmt.Println("c2 called")
	time.Sleep(5 * time.Second)
	fmt.Println("try to  call c1")
	c1()
	fmt.Println("c1 called")
	/*
		g1 start
		g3 start
		g2 start
		try to  call c3
		c3 called
		g3 done, err: <nil>
		try to  call c2
		c2 called
		g2 done, err: <nil>
		try to  call c1
		c1 called
		// 为什么 g1 done 没有出来？
	*/

}
