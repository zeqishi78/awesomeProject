package main

import "fmt"

func main() {
	//go语言提供了哪些集合类型的数据结构，数组、切片(slice)、map、list
	//go语言数组，var name [count]int
	var courses1 [3]string //courses类型，数组，只有3个元素的数据
	var courses2 [4]string
	fmt.Printf("%T\r\n", courses1) //[3]string
	fmt.Printf("%T\r\n", courses2) //[4]string
	// courses1 和 courses2 类型不一样
	courses1[0] = "go"
	courses1[1] = "grpc"
	courses1[2] = "gin"
	fmt.Println(courses1)
	//[]string和[3]string，是两种不同的类型，[]string是切片，[3]string是数组
	//如果元素个数是确定的不会变，那么声明一个指定容量的数组，是可以的，而且性能也很高，数组长度固定，性能比较高
	for _, value := range courses1 {
		fmt.Println(value)
	}

	//数组的初始化方式1
	var course3 [3]string = [3]string{"张三", "李四", "王五"}
	for _, value := range course3 {
		fmt.Println(value)
	}
	//数组初始化方式2
	course4 := [4]string{"数学", "语文", "英语", "政治"}
	for _, value := range course4 {
		fmt.Println(value)
	}
	//数组初始化方式3，在数组的指定位置放一个元素，其他位置不放元素，
	//那么默认值就是空
	course5 := [5]string{2: "湖北"}
	for _, value := range course5 {
		fmt.Println(value)
	}
	//数组初始化方式4，将数组长度使用省略号的方式
	course6 := [...]string{"北京", "上海", "广宗", "深圳"}
	for _, value := range course6 {
		fmt.Println(value)
	}
	for i := 0; i < len(course6); i++ {
		fmt.Println(course6[i])
	}
	course7 := [...]string{"北京", "上海", "武汉", "成都"}
	if course7 == course6 {
		fmt.Println("equas")
	}

	//多维数组
	var coursesInfo [4][4]string
	coursesInfo[0] = [4]string{"go", "1h", "bobby", "imooc"}
	coursesInfo[0][0] = "go"
	coursesInfo[0][1] = "1h"
	coursesInfo[0][2] = "bobby"
	coursesInfo[1] = [4]string{"grpc", "2h", "周扬", "尚硅谷"}
	coursesInfo[2] = [4]string{"gin", "3h", "胡大伟", "北京尚学堂"}
	coursesInfo[3] = [4]string{"docker", "4h", "nesta", "黑马程序员"}
	fmt.Println(len(coursesInfo))
	for i := 0; i < len(coursesInfo); i++ {
		for j := 0; j < len(coursesInfo[i]); j++ {
			fmt.Print(coursesInfo[i][j] + "  ")
		}
		fmt.Println()
	}
	//第二种遍历方式：
	fmt.Println("========第二种遍历方式=========")
	for _, value := range coursesInfo {
		fmt.Println(value)
	}
}
