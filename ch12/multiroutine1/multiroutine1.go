package main

import (
	"fmt"
	"time"
)

/*
	go func(){...}//创建并立即执行一个匿名函数作为Goroutinue，这个Goroutinue会独立于主线程运行，每秒打印bobby123
	1.Goroutinue是轻量级的线程，由go运行时管理。语法:go 函数名()，或者go func(){...}()，异步执行，不保证执行顺序
	2.闭包:闭包可以访问并操作器外部作用域的变量，在Goroutinue中使用闭包的时候，要特别注意变量捕获时机
	3.并发调度:Goroutinue的调度由Go运行时决定，创建Goroutinue的顺序不等于执行顺序，执行时受CPU核心数、调度策略影响
// 错误示例
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)  // 可能输出: 3, 3, 3
    }()
}

// 正确示例1：值传递
for i := 0; i < 3; i++ {
    go func(val int) {
        fmt.Println(val)  // 输出: 0, 1, 2（顺序不确定）
    }(i)
}

// 正确示例2：创建局部变量
for i := 0; i < 3; i++ {
    val := i  // 每次循环创建新变量
    go func() {
        fmt.Println(val)  // 输出: 0, 1, 2（顺序不确定）
    }()
}

*/

func asyncPrint() {
	for {
		time.Sleep(time.Second)
		fmt.Println("bobby12345")
	}
}
func main() {
	//匿名函数启动goroutinue
	go func() { //匿名函数,go声明goroutinue的过程，是go内部完成的，
		for {
			time.Sleep(time.Second)
			fmt.Println("bobby123")
		}
	}()
	//1.闭包
	//2.for循环的问题，for循环的时候，每个变量都会重用
	//每次for循环的时候，i变量都会被重用，当我们进行到第二轮的for循环的时候，这个i就发生变化了

	//闭包引用：匿名函数引用了外部变量的变量i
	//变量重用：i在每次循环中被重用
	//异步执行：goroutinue创建之后并不会立即执行，会等待调度
	//竞态条件：当Goroutinue实际执行的时候，循环可能已经结束，此时的i值已经是100
	//结果:可能打印出很多100，而不是期望的0-99
	for i := 0; i < 100; i++ {
		//每个循环都创建新的tmp变量，每个GoRoutinue引用自己的独立变量
		//解决方式一：创建局部变量，每次循环都是新的
		//tmp := i
		//go func() { //闭包，一个函数引用另外一个作用域的变量，但是在for循环的时候，每个变量会重用，goroutinue是一个异步的过程
		//	//不可能像同步的代码一样，有可能调度到最后一轮了，但是还没有调度到，先声明goroutinue并不代表这个goroutinue会先执行
		//	fmt.Println(tmp)	//引用局部变量
		//}()

		//第二种方式
		go func(i int) { //使用go关键字，启动一个goroutinue
			fmt.Println(i)
		}(i)

	}
	time.Sleep(10 * time.Second)
}
