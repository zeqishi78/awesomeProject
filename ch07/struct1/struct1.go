package main

import (
	"fmt"
	"strconv"
)

//type关键字
/*
	type关键字的用途
		(1)定义结构体
		(2)定义接口
		(3)定义类型别名
		type any = interface{}
		(4)类型定义
		(5)类型判断
*/
//别名实际上是为了更好的理解代码
type any = interface{}

func main() {
	//type MyInt = int //类型别名
	//var i MyInt = 12
	//var j int = 8
	//fmt.Printf("%T\r\n", i) //main.MyInt
	//fmt.Println(i + j)      //20，在编译的时候，类型别名，会被替换为int
	//
	type MyIntNew int //自定义类型，基于已有的类型，自定义一个类型
	var i1 MyIntNew = 12
	fmt.Printf("%T\r\n", i1) //main.MyIntNew

	var k MyInt = 12
	fmt.Println(k.string())

	//既希望你是int类型，又希望你可以增加方法
	//可以定义别名
	fmt.Println()

	var a interface{} = "abc"
	switch a.(type) {
	case string:
		fmt.Println("字符串")

	}
}

type MyInt int //自定义类型，基于已有的类型自定义一个类型
func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))
}
