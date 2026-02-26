package main

import "fmt"

/*
作用与设计

- 使用命名返回值 result、err，让 defer 闭包可以直接赋值 err 并改变函数最终返回。
- 在逻辑中对“除零”用 panic 表示严重运行时错误，交由 defer+recover 捕获并转换为 error 返回。
执行流程

- 进入 safeDivide，创建命名返回值 result=0、err=nil。
- 注册 defer 闭包；该闭包在函数返回或发生 panic 时执行。
- 判断 b 是否为 0：
  - 为 0：触发 panic("除零错误")，正常流程被中断，开始栈展开。
  - 不为 0：执行 return a/b, nil；随后 defer 仍会执行，但 recover() 返回 nil，不修改 err。
- 栈展开过程中执行 defer 闭包：
  - r := recover() 取出 panic 的值；若非 nil，说明发生了 panic。
  - 将 err 设为 fmt.Errorf("运行时错误: %v", r)，覆盖原来的 err（命名返回值生效）。
  - 函数结束返回：result 仍为默认 0（未修改），err 为包装后的错误。
- 在 main 中：
  - 调用 safeDivide(10, 0) 得到 err 非空，打印“捕获到错误: 运行时错误: 除零错误”。
  - 若调用 safeDivide(10, 2) 则返回 result=5，err=nil，打印“结果: 5”。
关键点

- recover 只能在发生 panic 的同一 goroutine 的 defer 中生效；
其他 goroutine 的 panic 需要在各自 goroutine 内用 defer+recover 处理。
- 使用命名返回值是为了让 defer 可以安全地修改返回的 err；这是 Go 处理“在退出前补充错误”的常用手法。
- 在无 panic 的正常路径下，recover 返回 nil，不影响返回值。
- panic 用于不可恢复的异常；像“除零这种可预期的业务错误”，更建议直接返回 error，而非 panic。
*/

func main() {
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("捕获到错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

}

/*
- safeDivide 使用命名返回值 result、err，并在函数内注册一个 defer 闭包用于捕获 panic 并将其转换为 error。
- main 调用 safeDivide(10, 0)，根据 err 是否为 nil 决定打印“捕获到错误”或“结果”。
*/
//进入到safeDevide方法，初始化命名返回值，result被初始化为0，err被初始化为nil
func safeDivide(a, b int) (result int, err error) {
	//执行defer注册语句，该函数会在safeDivide结束被调用，正常会遇到panic栈展开的时候，或者遇到return语句的时候
	defer func() {
		//r:=recover()捕获刚刚触发的panic值"除零错误"
		//r非nil，执行err=fmt.Errorf("运行时错误")
		//defer执行结束，safeDivide返回：result仍是默认值0，err是刚刚设置的错误值
		if r := recover(); r != nil {
			err = fmt.Errorf("运行时错误:%v", r)
		}
	}()
	//被除数为0，条件为真，执行panic，执行中断safeDivide的正常路径，后续的
	//return a/b,nil不会执行
	if b == 0 {
		//进入栈展开阶段，开始调用本函数内已经注册的defer
		panic("除零错误")
	}
	return a / b, nil
}
