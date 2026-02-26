package main

import "fmt"

/*
指针：提一个需求，我希望结构体，传值的时候，我在函数中修改的值，可以反应到变量中
*/
type Person struct {
	name string
}

func changeName(p *Person) {
	p.name = "imooc"
}

// 接收者
func (pn *Person) sayHello() {

}

// 通过指针交换两个值
func swap(a, b *int) {
	//临时值
	t := *a
	*a = *b
	*b = t
}

func main() {

	person := Person{
		name: "bobby",
	}
	changeName(&person)

	fmt.Println(person.name)
	//fmt.Println()
	var pi *Person = &person

	fmt.Printf("%p", pi)

	//指针的定义
	var pointer *Person
	pointer = &Person{}
	fmt.Println(pointer)
	/*
		第一个不同的点
		(1)go语言限制了指针的额运算，在c语言中你拿到了一个指针之后，进行+1的擦操作，在go语言中不行，不能参加运算
		(2)go的指针是一个阉割版，在unsafe包里面，并不是有bug，是要提醒你，这样做是不安全的额，一般不会使用unsafe包，但是你要使用的时候，是可以使用的
		(3)
		(4)
		(5)
	*/
	(*pointer).name = "booby5"
	fmt.Println(pointer)

	//var a int = 18
	//b := &a

	var c *int
	fmt.Println(c) //<nil>

	/*
		定义指针的时候，要初始化指针，否则会出现空指针
	*/
	//指针第一种初始化方式
	//ps := &Person{}
	//指针第二种初始化方式
	//var emptyPerson Person
	//pi2 := &emptyPerson

	//指针的第三种初始化方式
	var pp = new(Person)
	fmt.Println(pp.name)
	//map必须初始化

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
