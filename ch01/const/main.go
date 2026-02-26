package main

import "fmt"

func main() {
	//常量：定义的时候就指定的值，不能修改，常量尽量全部大写，
	const PI float32 = 3.141592654 //显式定义
	const PI2 = 3.141592654        //隐式定义
	////多个单词的常量，尽量使用下划线定义
	const MY_NAME = "张三"
	//PI = 3.2                       //常量不能修改

	const (
		UNKNOW = 1
		FEMALE = 2 //男女
		MALE   = 3
	)
	const (
		x int = 16
		y
		s = "abc"
		z
	)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(s)
	fmt.Println(z)
	/*
			常量类型只可以定义bool、数值(整数、浮点数和复数)和字符串
			不曾使用的常量，没有强制使用的要求，显式指定类型的时候，
		必须确保常量左右值类型一致
	*/

}
