package main

import "fmt"

func main() {
	//go语言接口，鸭子类型，php、python
	//go语言中，处处都是interface，到处都是鸭子类型duck typing
	/*
		当看到一只鸟，走起来像鸭子，游泳起来像鸭子，叫起来也像鸭子，那么这只鸟就是鸭子
		动词：方法，具备某些方法
	*/
	//创建pskDuck实例
	p1 := pskDuck{
		legs: 2,
	}
	gaga := p1.Gaga()
	fmt.Println(gaga)

	//创建指针类型
	duck2 := &pskDuck{
		legs: 4,
	}
	duck2.Walk()
	duck2.Gaga()

}

/*
	定义接口
	(1) 接受者类型决定接口实现。方法使用指针接受者*pskDuck，因此*pskDuck类型实现了duck接口
	(2)如果尝试将值类型赋值给duck类型接口变量，编译会报错
	(3)
	(4)
	(5)
	(6)

*/
//任何参数只要实现了这三个方法，就隐式满足了duck接口
type duck interface {
	//方法申请
	Gaga() string //返回字符串类型
	Walk()        //无参无返回值
	Swimming()    //无参无返回值
}

/*
定义了一个psDuck的结构体，包含了一个整型字段legs，表示鸭子腿数
*/
type pskDuck struct {
	legs int
}

/*
	下面三个方法都是为*pskDuck类型定义的，因此pskDuck的实例和指针都可以调用这些方法
*/
//接受者类型：*pskDuck指针类型

func (pd *pskDuck) Gaga() string {
	fmt.Println("嘎嘎")
	return "嘎嘎"
}

func (pd *pskDuck) Walk() {
	fmt.Println("this is pskDuck walking")
}

func (pd *pskDuck) Swimming() {
	fmt.Println("this is pskDuck Swimming")
}
