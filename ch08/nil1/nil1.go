package main

import (
	"fmt"
	"io"
)

/*
(1)go中nil是一个预定义的标识符，表示零值或者空值。
(2)nil表示无或者空，表示变量没有初始化，不指向任何有效内存值
(3)零值:对于指针、切片、map、channel、接口和函数类型，未初始化的时候，默认值就是nil
(4)
(5)
(6)
*/
func main() {
	//未初始化的指针为nil，解引用nil指针会导致panic
	var p1 *int
	fmt.Println(p1 == nil) //true
	var p2 *string
	fmt.Println(p2 == nil) //true

	//切片
	var s1 []int           //s1==nil
	fmt.Println(s1 == nil) //true
	s2 := []int{}          //s2!=nil，但是len=0,cap=0
	fmt.Println(s2 == nil) //s2!=nil，但是len=0，cap=0

	//map
	//nil map不能存储键值对，会panic，必须用make()初始化之后才能使用
	var m map[int]int     //m==nil
	fmt.Println(m == nil) //true

	//channel
	//nil channel发送和接收数据都会发生永久阻塞，必须用make()初始化
	var ch chan int        //ch ==nil
	fmt.Println(ch == nil) //true

	//接口
	var err error   //err==nil
	var w io.Writer //w==nil

	//接口包含(type,value)两部分，都为nil时，接口才为nil
	fmt.Println(err == nil) //true
	fmt.Println(w == nil)   //true

}
