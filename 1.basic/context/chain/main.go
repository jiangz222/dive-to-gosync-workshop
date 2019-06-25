package main

import (
	"context"
	"fmt"
)

/**
context.WithValue(ctx,k,v)将ctx和k、v重新封装成context后返回，
ctx.Value会轮询上面层层封装的ctx，直到找到指定的key为止
*/
func main() {
	ctx := context.Background()
	ctx = context.TODO()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0001")
	ctx = context.WithValue(ctx, "key3", "0001")
	ctx = context.WithValue(ctx, "key4", "0004")

	fmt.Println(ctx.Value("key1"))
}
