package main

import "fmt"

// 全局变量和局部变量
// var  name = "bobby"
// var age = 18
// var ok bool
var (
	name = "bobby"
	age  = 18
	ok   bool
)

func main() {
	//go是静态语言，静态语言和动态语言相比，变量差异很大
	//1.变量必须先定义，后使用
	//2.变量必须要有类型
	//3.类型定下来之后，不能改变
	//var name int
	//name =1
	var name1 int = 1
	age := 1
	//go语言中，变量定义了不使用是不行的
	fmt.Println(name1)
	fmt.Println(age)
	//多变量定义
	//var user1, user2, user3 string
	var user1, user2, user3 = "bobby1", 1, "bobby3"
	fmt.Println(user1, user2, user3)

	/*
		注意：变量必须先定义，才能使用
		go语言是静态语言，要求变量的类型和赋值类型一致
		变量名不能冲突
		**/
	var address string = "湖北省"
	fmt.Println(address)
	/*
			局部变量的变量名和全局变量的变量名可以重名，
			但是局部变量的优先级比全局变量的优先级高，
			如果打印的话，会优先打印局部变量
			变量是有0值的，如果定义了一个变量，定义了一个age2,如果不进行赋值，
			使用的时候，
		就不会像c++一样，
				**/
	fmt.Println(age)
	var age2 int
	fmt.Println(age2)

	/*
		如果定义了一个字符串，但是不进行赋值，打印的话，就是一个空字符串
		定义了变量，一定要使用，否则会报错
	**/
	var name3 string
	fmt.Println(name3)
	/*
		布尔类型的0值是false
	**/
	var ok2 bool
	fmt.Println(ok2)

}
