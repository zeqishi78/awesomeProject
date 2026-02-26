package main

import "fmt"

type Address struct {
	City    string
	Street  string
	ZipCode string
}

type Person struct {
	Name    string
	age     int
	Address Address
}

// 匿名嵌套
type Person1 struct {
	Name    string
	Age     int
	Address //匿名嵌套，字段提升
}

func main() {
	person := Person{
		Name: "张三",
		age:  12,
		Address: Address{
			"北京",
			"朝阳区",
			"10086",
		},
	}
	fmt.Println(person.Address.City)

	person1 := Person1{
		"李四",
		15,
		Address{
			City:    "北京",
			Street:  "朝阳区",
			ZipCode: "100000",
		},
	}
	//可以直接访问嵌入的字段
	fmt.Println(person1.City)    //等价于person1.Address.City
	fmt.Println(person1.ZipCode) //等价于person1.Address.ZipCdoe
}
