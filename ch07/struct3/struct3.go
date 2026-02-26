package main

import "fmt"

type Student3 struct {
	name string
	age  int
}

func (p Student3) prints() {
	//值传递，是拷贝副本，不会改变原来的值
	p.age = 35
	fmt.Printf("name : %s, age : %d", p.name, p.age)
}

func (p1 *Student3) prints1() {
	//指针传递，可以改变原来的值；或者数据较大的时候，考虑使用指针传递的方式
	p1.name = "王五"
	fmt.Printf("name : %s, age : %d", p1.name, p1.age)
}

//接收器有两种形态

func main() {
	student3 := &Student3{
		"张三",
		23,
	}
	student3.prints()
	student3.prints()
	student3.prints1()
}
