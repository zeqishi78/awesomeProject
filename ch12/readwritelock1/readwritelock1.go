package main

/*
1.读写锁的定义
	读写锁是一种特殊的互斥锁，允许多个读写操作同时进行，读写操作是互斥的
	import "sync"
	var rwLock sync.RWMutex
(1)普通锁sync.Mutex:无论是读还是写，同时只允许一个协程Goroutinue访问。如果有很多读操作，他们必须排队，效率较低
(2)读写锁sync.RMutex
	①读锁Rlock：允许多个协程同时持有读锁。这意味着如果有10个协程同时读取数据，他们可以并行执行，互补阻塞。
	②写锁Lock：是独占的。当有一个协程持有写锁时，其他所有想读、写的协程都必须等待。
	③互斥规则:
		读 + 读:允许并发
		读 + 写:互斥(不能同时发生)
		写 + 写:互斥(不能同时发生)
	④适用场景:非常适合读取频率远高于写入频率的场景
(3)
(4)
(5)
(6)

2.与普通Mutex对比
	特性			sync.Mutex		sync.RMutex
	读锁			与写锁相同		可共享，允许多个读
	写锁			互斥				互斥
	性能			读写频繁			读多写少

3.核心方法
	//读锁方法
	//获取读锁(可共享)
	rwLock.RLock()
	//释放读锁
	rwLock.RUnlock()

	//写锁方法
	//获取写锁(互斥)
	rwLock.Lock()
	//释放写锁
	rwLock.Unlock()



*/
import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex            //全局的读写锁实例，用于保护下面的data
var data = make(map[string]string) //data:一个字符串映射，这是共享资源。在go中，map本身不是并发安全的，如果多个协程同时读写同一个map，并且不加锁，程序会直接发生panic

/*
逻辑:调用Lock()获取写锁。此时，如果有其他协程正在读或者正在写，当前协程会阻塞，直到锁可用
安全性:一旦获取锁，当前协程拥有对data的完全控制权，其他任何读写操作都被挡在门外
defer:使用defer是为了防止如果在临界区内发生panic，锁也能被自动释放，避免死锁
*/
func write(key, value string) {
	rwLock.Lock()         //获取写锁,独占
	defer rwLock.Unlock() //确保释放，确保函数退出前释放锁
	fmt.Printf("Writing : %s = %s\r\n", key, value)
	time.Sleep(100 * time.Millisecond) //模拟写耗时
	data[key] = value                  //修改共享数据
	fmt.Printf("write complete: %s = %s\r\n", key, value)
}

/*
逻辑:调用RLock()获取读锁
并发性:

	①如果没有写锁，多个read协程可以同时通过这一行，并行执行后面的Sleep和读取操作
	②如果当前有写锁(write函数正在运行)，所有的read写成都会在这里阻塞，等待操作完成
*/
func read(key string) string {
	rwLock.RLock() //获取读锁
	defer rwLock.RUnlock()
	fmt.Printf("Reading: %s\r\n", key)
	time.Sleep(50 * time.Millisecond) //模拟读取耗时
	value := data[key]
	fmt.Printf("Read complete: %s = %s \r\n", key, value)
	return value
}

/*
WaitGroup：这是一个计数器，确保main函数不会在子协程运行完成之前退出。
并发竞争:

	(1)程序启动了11个协程组成，10读 + 1写
	(2)由于Go调度器的不确定性，谁先拿到锁时不确定的
	(3)情况A读优先:如果10个协程先运行，他们会全部获得RLock，并行打印Reading...，同时休眠50ms。
		此时写协程Lock()会被阻塞，直到10个读协程全部调用RUnlock()。
	(4)情况B写优先:如果写协程先运行，它获得Lock。此时所有10个读协程调用RLock()都会被阻塞。写协程
		休眠100ms，完成后10个协程才能依次或者批量获得读锁。
	(5)情况C(混合):部分读协程先执行，写协程插入等待，或者写协程执行一半时的读请求到来，会被则是
	(6)代码执行时的预期现象:
		①读的并发性:如果看到多个Reading:key...几乎同时打印出来。因为读锁是共享的，他们不需要互相等待
		②写的独占性:Writing和write complete之间，不会穿插任何Reading和Read complete输出。
			因为写锁持有期间，禁止任何读操作进入临界区。
		③
		④
		⑤
		⑥
*/
func main() {
	var wg sync.WaitGroup //用于等待所有的协程结束

	//启动多个读协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			read(fmt.Sprintf("key%d", id%2))
		}(i)
	}

	//启动写协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		write("key1", "new_value")
	}()
	wg.Wait()
}
