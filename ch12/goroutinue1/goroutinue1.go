package main

import (
	"fmt"
	"sync"
)

// 子goroutinue如何通知到主goroutinue自己技术了，主goroutinue如何知道自子goroutinue已经结束了？
// waitGroup主要用于goroutinue的执行等待，Add方法要和Done方法配套
func main() {
	var wg sync.WaitGroup
	//我要监控多少个goroutinue执行结束
	//如果i是固定的，可以在for循环外面写
	// Add和Done一定要成对出现
	wg.Add(100)
	for i := 0; i < 100; i++ {
		//如果100是一个变量，可以在内部这样写
		//wg.Add(1)
		go func(i int) {
			//defer是为了防止大家忘掉，所以这里加一个defer，保证整个函数执行完成之后再执行
			defer wg.Done() //调用一次Done之后，就会把Add(100)中的100减去1，全部执行完了之后，然后退出
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("all done")
}
