package main

import "fmt"

/*
指针与结构体
*/

type Person struct {
	Name string
	Age  int
}

/*
birthday函数接收一个指向Person类型的指针参数p
*/
func birthday(p *Person) {
	//通过指针修改结构体字段：将p指向Person的Age字段加1
	//go语言中，(*p).Age++等价于p.Age++
	(*p).Age++
}

func main() {
	//创建结构体指针
	//使用&Person{}创建Person实例并返回其指针
	person := &Person{
		Name: "张三", Age: 34,
	}
	fmt.Printf("生日前:%s今年%d岁\r\n", person.Name, person.Age)
	birthday(person)
	fmt.Printf("生日后:%s今年%d岁\r\n", person.Name, person.Age)

	//直接通过指针修改字段
	person.Name = "胡思"
	fmt.Printf("改名后:%s\r\n", person.Name)
}
