package main

import (
	"errors"
	"fmt"
)

func main() {
	//1.调用多返回值函数
	f, err := divide(3.14, 1.1)
	if err != nil {
		fmt.Println("函数报错了...")
	} else {
		fmt.Println(f)
	}
	getDimensions()
}

//基本函数声明
/*
func functionName(para1 type1) returnType{
	//函数体
	return value
}
*/

// 多返回值
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 命名返回值
func getDimensions() (width int, height int) {
	width = 200
	height = 300
	return
}

/*
	1.每个函数都有唯一的签名，由参数类型和返回值类型决定

*/
//以下两个函数类型相同
func add(x, y int) int {
	return x + y
}
func substract(x, y int) int {
	return x - y
}

// 以下函数类型不同
func multiply(x, y float64) float64 {
	return x * y
}
