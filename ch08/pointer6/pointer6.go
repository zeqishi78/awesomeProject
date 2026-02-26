package main

import "fmt"

/*
使用new()函数动态分配内存
(1)语法：new(Type)，返回指针*Type
(2)功能：在堆上分配指定类型的内存空间，并返回指向该内存的指针
(3)初始化：分配的内存会被初始化为该类型的零值

内存分配对比
(1)var p int：在栈上分配int变量，零值初始化
(2)p :=new(int)：表示在堆上分配int内存，返回指针，零值初始化

结构体指针
(1)new(point)：返回*point类型的指针
(2)可以通过指针直接访问字段：pointPtr.x等同于(*pointPter).x
*/
type point struct {
	X, Y int
}

func main() {
	//使用new分配int类型内存
	//new(int)在堆上分配int大小的内存，返回指向该内存的指针
	//内存会被初始化为零值
	ptr := new(int)
	//通过指针赋值，将分配的内存设置为100
	*ptr = 100
	//%p将指针格式化为十六进制表示
	fmt.Printf("指针地址:%p\r\n", ptr)
	//打印指针指向的值
	fmt.Printf("指针值:%d\r\n", *ptr)

	//使用New分配结构体内存
	//new(Point)，表示在堆上分配Point结构体大小的内存
	//所有字段都会被初始化为零值，int字段为0
	pointPtr := new(point)
	pointPtr.X = 10
	pointPtr.Y = 20
	fmt.Printf("点坐标:(%d,%d)\r\n", pointPtr.X, pointPtr.Y)
}
