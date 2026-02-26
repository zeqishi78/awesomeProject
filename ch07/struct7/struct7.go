package main

import "fmt"

type Rectangle struct {
	Width  float32
	Height float32
}

// 为Rectanble定义方法
func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

// 指针接收者方法，可以修改结构体
func (r *Rectangle) Scale(factor float32) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	fmt.Println("面积:", rectangle.Area())

	rectangle.Scale(2)
	fmt.Println("缩放后:", rectangle.Width, rectangle.Height)
}
