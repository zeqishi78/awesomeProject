package main

import "fmt"

type Person struct {
	Name    string
	age     int
	Eamil   string
	Address string
}

// 创建与初始化
func main() {
	//方式一：使用字段名初始化
	person1 := Person{
		Name:    "张三",
		age:     23,
		Eamil:   "1528129110@qq.com",
		Address: "广东省广州市越秀区",
	}
	fmt.Println(person1)
	//方式二：顺序初始化
	person2 := Person{
		"李四", 34, "892549346@qq.com", "湖北省武汉市",
	}
	fmt.Println(person2)
	//方式三：不分初始化，未初始化的字段位零
	person3 := Person{
		Address: "江苏省南京市鼓楼区",
	}
	fmt.Println(person3)
	//方式四：先声明后赋值
	var person4 Person
	person4.Name = "赵六"
	person4.age = 36
	fmt.Println(person4)

}
