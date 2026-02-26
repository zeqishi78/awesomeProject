package main

import "fmt"

/*
我现在想要放多个person的信息，到一个集合中，
*/
//类型集合
type Person struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {
	//如何初始化结构体
	p1 := Person{"bobby", 12, "mooc", 180.2}
	fmt.Println(p1)
	p2 := Person{"张三", 18, "湖北省", 175.5}
	fmt.Println(p2)
	//第二种初始化方式]
	p3 := Person{
		name: "张三",
		age:  23,
	}
	fmt.Println(p3)

	var persons []Person
	persons = append(persons, p1)

	p4 := []Person{
		{"张三", 23, "湖北省", 23.34},
		{"李四", 24, "湖南省", 25.56},
		{"王五", 25, "河南省", 26.67},
		{"赵六", 26, "河北省", 27.78},
		{"胡七", 27, "上海市", 28.76},
	}
	fmt.Println(p4)

	var p5 Person
	p5.name = "胡小三"
	p5.age = 23
	p5.address = "湖北省"
	p5.height = 185.5
	fmt.Println(p5.name)

	//匿名结构体，类似于匿名函数
	add := struct {
		province string
		city     string
		address  string
	}{
		"北京市",
		"通州区",
		"九资河",
	}
	fmt.Println(add)
}
