package main

import (
	"fmt"
	"time"
)

// go函数支持普通函数、匿名函数、闭包
func main() {
	/*
		go中函数是一等公民
			(1)函数本身可以当做变量
			(2)函数包含了匿名函数和闭包
			(3)函数可以满足接口
	*/
	fmt.Println(add(2, 3))
	sum, _ := add(3, 4)
	fmt.Println(sum)
	a := 1
	b := 2
	sum1, _ := add1(a, b, 3.13)
	fmt.Println(sum1)
	sum2, _ := add1(a, b, 5)
	fmt.Println(sum2)
	fmt.Println("=======================================")
	fmt.Printf("main函数内部打印：a= %d,b=%d\r\n", a, b)
	fmt.Println(add3(1, 2, 3))

}

// go中可以返回多值的
func add(a, b int) (int, error) {
	return a + b, nil
}

// 函数参数传递的时候，到底是值传递还是引用传递？答：go语言中，全部都是值传递
//
// 如果参数类型不一致，就不能省略参数类型
func add1(a, b int, c float32) (float32, error) {
	a = a + 1
	b = b + 2
	fmt.Printf("add1函数内部打印：a= %d,b=%d", a, b)
	return float32(a) + float32(b) + c, nil
}

// 返回值，除了类型之外，还可以显式指定类型变量名
func add2(a, b int, c float32) (f float32, err error) {
	sum := float32(a) + float32(b) + c
	return sum, err
}

// 省略号，可变参数，变长类型，要求类型都相同
func add3(item ...int) int {
	sum := 0
	for _, value := range item {
		sum += value
	}
	return sum
}

// 无返回值的参数
func runForever() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("doing...")
	}
}
