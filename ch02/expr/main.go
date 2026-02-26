package main

import "fmt"

func main() {
	//运算符 + - * / % ++ --
	var a, b = 1, 2
	var astr, bstr = "hello", "bobby"
	fmt.Println(a, b, a+b)
	fmt.Println(astr, bstr, astr+bstr)

	fmt.Println(3 % 2)
	a++
	fmt.Println(a)

	//逻辑运算符
	var abool, bbool = true, false
	if abool || bbool {

	}
	//位运算符
	var A = 60
	var B = 13
	fmt.Println(A & B) //12

	d := &A //取地址
	fmt.Println(d)
}
