package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	//上面类型的字段不用重复定义
	//第一种嵌套方式
	p     Person
	score float32
}

// 第二种嵌套方式，匿名嵌套
type Student1 struct {
	Person
}

func main() {
	//结尾要带,
	//第一种嵌套方式
	s := Student{
		Person{
			"BOBBY",
			18,
		},
		98.8,
	}
	name := s.p.name
	s.p.name = "张三"
	fmt.Println(name)
	s1 := Student1{
		Person{"张三", 23},
	}
	fmt.Println(s1.name)

}
