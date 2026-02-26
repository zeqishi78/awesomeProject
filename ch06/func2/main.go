package main

import "fmt"

/*
go函数是一等公民
1.函数本身可以当做变量
2.匿名函数  闭包
3.函数可以满足接口
*/
func main() {
	//传递函数名，不能传递调用
	funVar := add1
	funVar(1, 2)
	cal("+", 1, 2, 3)()
	sum := cal1(func(items ...int) int {
		sum := 0
		for _, v := range items {
			sum += v
		}
		return sum
	})
	fmt.Println(sum) //0
	callback(1, add2)
}

func add1(a, b int) (int, error) {
	sum := a + b
	return sum, nil
}

func add2(a, b int) {
	fmt.Printf("sum is %d", a+b)
}

// 返回值为函数的情况
func cal(op string, items ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("这既不是加法，也不是减法")
		}
	}
}

// myFunc返回一个int
func cal1(myFunc func(items ...int) int) int {
	return myFunc()
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
