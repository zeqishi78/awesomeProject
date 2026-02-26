package main

import "fmt"

/*
nil指针：
(1)声明但是没有初始化的指针，默认值为nil
(2)var ptr *int中的ptr初始化值为nil
(3)nil指针不指向任何有效内存地址

指针安全检测
(1)解引用nil指针会导致运行时panic
(2)必须在试用前检查指针是否为nil，if ptr ! = nil

指针有效性
(1)
*/
func main() {
	//声明一个int类型指针ptr，未初始化，默认值为nil
	var ptr *int
	//测试nil指针:调用函数传入nil指针
	printIfValid(ptr)
	//声明并初始化整型变量num，值为50
	num := 50
	//把num的地址赋值给ptr，ptr现在指向有效的内存地址
	ptr = &num
	printIfValid(ptr)

	var ptr1 *int
	fmt.Println(ptr1) //<nil>

	//解引用nil指针
	fmt.Println(*ptr1) //panic: runtime error: invalid memory address or nil pointer dereference
}

/*
printIfValid函数接收一个int类型指针参数
*/
func printIfValid(ptr *int) {
	//指针有效，可以安全解引用
	if ptr != nil {
		fmt.Printf("有效值，值:%d\r\n", *ptr)
	} else {
		//指针为nil，不能解引用，否则会引发panic
		fmt.Println("nil指针，无法解引用")
	}
}
