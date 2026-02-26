package main

import "fmt"

func main() {
	//var a int8
	//var b int16
	//var c int32
	//var d int64
	//
	//var ua uint8
	//var ub uint16LL
	//var uc uint32
	//var ud uint64
	//
	//var e int //动态类型，用的时候就会知道，用起来挺麻烦的
	//
	//a = int8(b) //类型转换，必须指定类型，如果强转的话，精度可能会丢失
	//
	////float类型
	//var f1 float32 //10的32次方
	//var f2 float64

	//var c byte
	//用于存放字符的，go语言中，没有类似java中的char类型，
	//主要用来存放字符，byte就是type byte = uint8
	//c = 'a'
	//fmt.Println(c)

	var c byte //主要适用于存放字符
	c = 'a' + 5
	fmt.Printf("c=%c", c)
	fmt.Println()
	c1 := 97
	fmt.Printf("c=%c", c1)

	var c2 rune
	c2 = '张'
	fmt.Printf("c=%c", c2)

	var name string
	name = "张三"
	fmt.Println(name)

}
