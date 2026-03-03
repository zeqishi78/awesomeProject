package main

import (
	"fmt"
	"sync"
	"time"
)

/*
锁本质上，是将并行的代码串行化，使用lock肯定会影响性能，
即使是设计锁，那么应该尽量保证并行
我们有两组协程，其中一组负责写数据，另一组负责读取数据，web系统中，绝大部分场景都是读多写少，甚至是读远远多于写，
虽然有多个goroutinue，但是仔细分析，我们会发现，读协程之间应该并发，读和写之间应该串行，当我在写的时候，你是不能读取数据的
读和读之间，也不应该并行
*/
//读写锁

func main() {
	var rwlock sync.RWMutex
	var num int

	var wg sync.WaitGroup
	wg.Add(2)

	//负责写数据
	go func() { //写数据goroutinue
		defer wg.Done()
		rwlock.Lock() //加写锁，写锁会防止别的写锁获取，和读锁获取
		num = 12
		defer rwlock.Unlock()
		time.Sleep(time.Second * 5)
		fmt.Println("...get write lock")
	}()
	time.Sleep(time.Second)

	//负责读数据
	go func() { //读的goroutinue
		wg.Done()
		for {
			rwlock.RLock() //加读锁，读锁不会阻止别人的读取
			defer rwlock.RUnlock()
			time.Sleep(500 * time.Millisecond)
			fmt.Println(num)
			fmt.Println("get read lock")
		}

	}()
	wg.Wait()
}
