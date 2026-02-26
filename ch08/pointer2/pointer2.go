package main

import "fmt"

func main() {
	//声明并初始化一个整形变量num，赋值为100
	var num int = 100
	//声明一个指向int类型的指针变量ptr，此时ptr为nil(空指针)
	var ptr *int
	//1.将num的地址赋值给ptr
	//&num是获取num的内存地址，ptr现在指向num
	ptr = &num
	fmt.Println(num)
	fmt.Println(ptr)

	//2.通过指针修改num的值
	//*ptr解引用，将ptr指向的地址改为200
	//相当于直接修改num=200
	*ptr = 200

	//3.输出结果验证
	fmt.Printf("num的值:%d\r\n", num)
	fmt.Printf("ptr指向的值:\r\n", *ptr)
	fmt.Printf("ptr的地址:%p\r\n", ptr)
}
