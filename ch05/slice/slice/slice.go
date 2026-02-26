package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
值传递 (Pass by Value)

传递什么：传递变量值的副本

核心特性：方法内对参数的修改不会影响原始变量

类比：给你一份文档的复印件，你在复印件上修改，不影响原件

引用传递 (Pass by Reference)

传递什么：传递变量自身的引用/别名

核心特性：方法内对参数的修改直接影响原始变量

类比：给你原文档的编辑权限，你的修改直接影响原件
*/
func printSlice(data []string) {
	data[0] = "java"
	//for i := 0; i < 10; i++ {
	//	data = append(data, strconv.Itoa(i))
	//}
}

type slice struct {
	array unsafe.Pointer //用来存储实际数据的数组指针，指向一块连续的内存
	len   int            //切片中元素的数量
	cap   int            //array数组的长度
}

func main() {
	//go的切片slice，在函数参数传递的时候，是值传递？还是引用传递？
	//答：严格来说是值传递，但是效果的话，又呈现出来了引用传递的效果，说它是引用传递的话，又不完全是引用传递
	//在声明一个slice的时候

	courses1 := []string{"go", "grpc", "gin"}

	printSlice(courses1)
	//如果是值传递的话，那我在外面改的值，那你在里面不应该受到影响
	fmt.Println(courses1)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[1:6]
	s2 := data[2:7]
	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 123, 1, 2, 3, 4, 5, 6, 7)

	s2[0] = 22
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	var courses2 []string
	for i := 0; i < 2000; i++ {
		//courses1 = append(courses2, strconv.Itoa(i))
		courses2 = append(courses2, strconv.Itoa(i))
		fmt.Printf("len: %d,cap: %d\r\n", len(courses2), cap(courses2))
	}

}
