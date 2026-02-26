package main

import (
	"fmt"
	"math"
)

/*
	1.接口是一种抽象类型，它定义了一组方法的签名，包括方法名、参数和返回值，但是不包含方法的视线，接口是go语言视线多台和鸭子类型的核心机制
	2.接口定义
		type 接口名 interface {
			方法名1 (参数列表) 返回值列表
			方法名2(参数列表) 返回值列表
		}
*/
//接口定义
type Writer interface {
	Writer([]byte) (int, error)
}

/*
接口实现(隐式实现)
(1)go语言中，接口实现是隐式的，不需要像java那样显式的使用implements关键字显式声明，只要一个类型实现了接口中的所有方法，它就自动满足该接口
(2)
()
()
()
()
*/
//定义一个结合形状的接口
//接口定义：Shape接口要求实现Area()方法和Perimeter()方法
//go语言中，类型不需要显式声明实现，只要实现了接口的所有方法就自动满足该接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 矩形实现Shape接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 圆形结构体
type Circle struct {
	Radius float64
}

// 圆形实现Shape接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 打印新装信息
func printShapInfo(s Shape) {
	fmt.Printf("形状信息:面积 = %.2f,周长 = %2.f\r\n", s.Area(), s.Perimeter())
}

// 空接口可以接收任何类型
// 空接口：interface{}表示可以接收任何类型的值
// 用途：用于实现类似泛型的功能，处理未知类型的数据
func describe(i interface{}) {
	fmt.Printf("值:%d,类型:%T\r\n", i, i)
}

// 类型断言示例
// 功能：检查接口值的实际类型，并安全的转换
func processEmptyInterface(i interface{}) {
	//类型断言
	if str, ok := i.(string); ok {
		fmt.Printf("这是字符串%s\r\n", str)
	} else if i, ok := i.(int); ok {
		fmt.Printf("这是数字:%d\r\n", i)
	} else {
		fmt.Printf("未知类型:%T\r\n", i)
	}
}

// 类型开关
func processSwithTypeSwitch(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("字符串长度:%d\r\n", len(v))
	case int:
		fmt.Printf("数字的平方:%d\r\n", v*v)
	case Shape:
		fmt.Printf("这是形状，面积:%f2.f\r\n", v.Area())
	default:
		fmt.Printf("未处理的类型:%T\r\n", v)
	}
}

func main() {
	r := Rectangle{
		Height: 3,
		Width:  4,
	}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())
	c := Circle{
		Radius: 3,
	}
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())

	describe(r)
	describe(c)
	processEmptyInterface(r)
	processEmptyInterface(c)

	processSwithTypeSwitch(r)
	processSwithTypeSwitch(c)
}
