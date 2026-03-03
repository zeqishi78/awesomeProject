package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
渐进式的方式
goroutinue监控cpu的信息
*/

var wg sync.WaitGroup

// 我们有一个新的需求，我们可以主动退出监控程序
// 共享变量

func cpuInfo(ctx context.Context) {
	fmt.Sprintf("traceid:%s\r\n", ctx.Value("traceid"))
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu信息")
		}
	}
}

func main() {
	//渐进式的方式
	wg.Add(1)
	//context包提供了三种结函数，withCancle，withTimeout，withValue
	//如果你的Goutinue函数中，如果希望被控制，超时消息、传值消息，但是我不希望影响原来接口信息的时候，函数参数中，第一个参数就尽量要加上ctx，
	//ctx1, cancelFunc1 := context.WithCancel(context.Background())
	//ctx2, _ := context.WithCancel(ctx1) //超时需求
	//2.主动超时  timeout
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//WithDeadLine在时间点cancle

	//withValue
	valueCtx := context.WithValue(ctx, "traceId", "gfasd")
	go cpuInfo(valueCtx)
	wg.Wait()
	fmt.Println("监控完成...")
}
