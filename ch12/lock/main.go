package main

import (
	"fmt"
	"sync"
)

//多个goroutinue，会对共享变量进行操作
/*
	锁就	- 资源竞争
	下面代码会启动两个Goroutinue：
		1.add():对total变量加1，执行10万次
		2.sub():对total变量减1，执行10万次
	理论上，最终的total值是0，但是实际上，如果没有锁的保护，结果通常不会是0，而是一个随机数
*/

var total int         //共享变量
var wg sync.WaitGroup //等待组，用于等待所有GoRoutinue完成
var lock sync.Mutex   //互斥锁，保护共享变量

func operation() {
	lock.Lock()
	//临界区代码
	lock.Unlock() //解锁
}

/*
锁本身是可以复制的，但是注意，锁千万不要复制，复制之后就失去了锁的效果了，锁的本质就是不停的改变锁的状态
*/
func add() {
	defer wg.Done() //函数结束的时候通知 WaitGroup
	for i := 0; i < 100000; i++ {
		lock.Lock()   //获取锁
		total += 1    //临界区操作
		lock.Unlock() //释放锁
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()   //获取锁
		total -= 1    //临界区操作
		lock.Unlock() //释放锁
	}
}

/*
这段代码的目的:

	(1)启动两个Goroutinue同时执行
	(2)主线程等待两个字Goroutinue都执行完毕
	(3)确保add和sub都完成后，再打印最终结果
	(4)wg.Add(2)作用是初始化等待组的计数器为2；WaitGroup内部有一个计数器，wg.Add(n)：计数器增加n，wg.Done()：计数器减1；wg.wait()：阻塞，直到计数器归零
	(5)wg.Wait():阻塞，直到计数器归零，必须在启动Goroutinue之前调用wg.Add()，否则可能出现竞态条件
	(6)go add()和go sub()。异步执行：这两个函数会立即执行，不会阻塞主线程；并发执行:add()和sub()会同时运行；
	(7)wg.Wait():等待两个Goroutinue。阻塞：主线程会在这里暂停执行，等待计数器归零；唤醒条件：当两个Goutinue都调用wg.Done()之后，
*/
func main() {
	wg.Add(2)          //设置需要等待2个Goroutinue
	go add()           //启动加法Goroutinue
	go sub()           //启动减法Goroutinue
	wg.Wait()          //等待两个Goroutinue都完成
	fmt.Println(total) //输出最终结果
}

/*
下面这部分代码存在资源竞争，
total+=1不是一条原子操作，实际上对应三条机器指令
(1)从内存读取total到寄存器
(2)寄存器值加1
(3)将寄存器值写回到内存

func add()  {
	for i := 0; i < 100000; i++ {
		total+=1
	}
}
时间线
初始值：total = 0

时间线：
t1: Goroutine1 读取 total = 0
t2: Goroutine2 读取 total = 0
t3: Goroutine1 计算 0+1 = 1
t4: Goroutine2 计算 0-1 = -1
t5: Goroutine1 写入 total = 1
t6: Goroutine2 写入 total = -1

最终结果：total = -1（应该是 0）

互斥锁
锁的特性：
	互斥锁：同一时间只有一个Goroutinue能获得锁
	阻塞性：其他Goroutinue尝试加锁时回阻塞等待
	可重入性：Go的Mutex是不可重入



*/
