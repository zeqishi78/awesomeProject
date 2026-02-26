package main

import "fmt"

func main() {
	//定义函数类型变量
	var operation func(int, int) int

	//赋值
	operation = func(a, b int) int {
		return a + b
	}
	fmt.Println(operation(2, 3))
	/*
		重新赋值
	*/
	operation = func(i, j int) int {
		return i * j
	}
	fmt.Println(operation(5, 6))
	//使用
	numbers := processNumbers([]int{1, 2, 3}, func(x int) int {
		return x * x
	})

	for _, v := range numbers {
		fmt.Println(v)
	}

}

// 函数作为参数
func processNumbers(numbers []int, transformer func(int) int) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = transformer(n)
	}
	return result
}
