package main

import "fmt"

/*
go语言闭包
我希望有个函数，每次调用一次，返回的结果值，
都是增加一次之后的值
*/
func main() {
	counter := markCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	//创建另一个独立的计数器
	counter1 := markCounter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	//闭包特性1：状态保持
	multiplier := makeMultiplier(2)
	fmt.Println(multiplier(5))
	fmt.Println(multiplier(3))

}

func makeMultiplier(i int) func(int) int {
	return func(j int) int {
		return i * j
	}
}

func markCounter() func() int {
	count := 0 //外部变量
	return func() int {
		count++ //引用并修改外部变量
		return count
	}
}
