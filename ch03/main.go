package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	name := "imooc体系课程学习"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	//转义符
	courseName := "go\"体系课\""
	fmt.Println(courseName)

	courseName1 := `go"体系课"`
	fmt.Println(courseName1)

	courseName2 := "go\r\n体系课"
	fmt.Println(courseName2)

	fmt.Print("hello")
	fmt.Println("world")

	//格式化输出
	uname := "bobby"
	out := "hello  " + uname
	fmt.Println(out)

	//格式化输出
	username := "张三"
	age := 18
	address := "北京"
	mobile := "1899999999"
	//拼凑过程比较麻烦，比较难以维护
	fmt.Println("用户名 : "+username, ", 年龄 :"+strconv.Itoa(age), ", 地址 : "+address, " , 手机号 : "+mobile) //极其难以维护
	//这个会非常常用，但是性能没有ln性能好，
	fmt.Printf("用户名:%s,年龄:%d,地址:%s,手机号:%s\r\n", username, age, address, mobile) //Printf不会加回车换行符，需要在最后的%s的地方加一个回车换行符

	userMsg := fmt.Sprintf("用户名:%s,年龄:%d,地址:%s,手机号:%s\r\n", username, age, address, mobile)
	fmt.Println(userMsg)

	//%v打印格式
	var ages []int = []int{1, 2, 3}
	fmt.Printf("ages:%v\r\n", ages) //会把原本的格式输出
	//%T：输出数据类型
	fmt.Printf("ages的类型: %T，name的类型:%T \r\n", ages, name) //[]int

	//通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(username)
	builder.WriteString(",年龄:")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(",地址:")
	builder.WriteString(address)
	builder.WriteString(",手机号:")
	builder.WriteString(mobile)
	fmt.Println(builder.String())
}
