package main

import "fmt"

/*
执行顺序

- main 注册 defer “main 的 defer”
- main 打印“开始”
- 调用 testPanic
- testPanic 注册 defer “这个 defer 会执行”
- testPanic 触发 panic，中断正常流程
- 在 testPanic 返回栈展开时，执行其 defer，打印“这个 defer 会执行”
- 回到 main 的栈展开，执行 main 的 defer，打印“main 的 defer”
- 程序未恢复 panic，运行时打印 panic 信息并终止；main 中的“结束”不会执行
打印结果

- 开始
- 这个 defer 会执行
- main 的 defer
- panic: 发生严重错误
关键点

- defer 在函数正常返回或发生 panic 时都会执行，执行顺序为后注册先执行（LIFO）
- panic 触发后，当前函数后续语句（如“这行不会执行”）不会再执行
- 若希望程序在 panic 后继续，可在上层用 defer+recover 处理并恢复
recover 示例

- 在 main 顶层捕获并恢复，确保“结束”可以执行
*/
func main() {
	/*
		执行顺序：
			(1)开始
			(2)这个defer会执行
			(3)main的defer
			(4)panic: 发生严重错误
	*/
	defer fmt.Println("main的defer")

	fmt.Println("开始")
	testPanic()
	fmt.Println("结束")

}

/*
defer在函数正常返回，或者发生panic时都会执行，执行顺序为先注册后执行LIFO
panic触发后，当前函数后续语句不会再执行
*/
func testPanic() {
	defer fmt.Println("这个defer会执行")
	panic("发生严重错误")
	fmt.Println("这行不会执行")
}
