package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	/*
		不同数据类型的零值是不一样的
		bool	false
		numbers	0
		string	""
		pointer	nil
		slice	nil
		map		nil
		channel,interface,function	nil
		struct默认值不是nil，默认值是具体字段的默认值
	*/
	person1 := Person{
		name: "bobby3",
		age:  12,
	}

	person2 := Person{
		name: "bobby5",
		age:  23,
	}
	if person1 == person2 {
		fmt.Println("yes")
	}

	//slice的默认值的问题
	var ps []Person //nil的slice
	if ps == nil {
		fmt.Println("nil slice")
	}
	var ps2 = make([]Person, 0) //empty的slice
	if ps2 == nil {
		fmt.Println("nil slice")
	}

	var m map[string]string //nil的map
	if m == nil {
		fmt.Println("nil slice")
	}

	var m2 = make(map[string]string) //empty的map
	if m2 == nil {
		fmt.Println("nil slice")
	}

}
