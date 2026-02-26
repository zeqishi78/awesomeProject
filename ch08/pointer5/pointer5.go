package main

import "fmt"

/*
指针数组与切片
(1)指针切片存储的是内存地址，而不是值的副本
(2)修改指针指向的值会影响原始数据
(3)如果原始切片扩容重新分配内存，指针可能会失效
(4)
(5)
(6)
(7)
(8)
*/
func main() {
	//创建整数类型切片numbers，包含5个元素
	numbers := []int{1, 2, 3, 4, 5}

	//声明一个指向int类型的指针切片
	//声明一个存储int指针的切片
	var pointers []*int
	//遍历numbers切片，获取每个元素地址
	for i := range numbers {
		//将每个元素地址添加到points切片中
		//&numbers[i]表示获取第i个元素的内存地址
		//&numbers[i]获取切片中指定元素的地址
		pointers = append(pointers, &numbers[i])
	}
	//打印
	fmt.Println("pointers = ", pointers)

	fmt.Println("修改前的数组:", numbers)
	//通过指针修改原来数组的值
	for _, ptr := range pointers {
		//*ptr解引用获取指针指向的那个值，然后乘以2
		//相当于直接修改numbers数组中对应的元素
		*ptr *= 2
	}
	//输出修改后的原数组
	fmt.Println("修改后的数组:", numbers)
	//输出指针指向的值
	for i, ptr := range pointers {
		fmt.Printf("pointers[%d]指向的值:%d\r\n", i, *ptr)
	}
}
