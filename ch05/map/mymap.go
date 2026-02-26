package main

import (
	"fmt"
)

func main() {
	//map是一个key(索引)和value(值)的无序集合，主要是查询方便
	//var coursesMap map[string]string
	//初始化方式1：初始化赋值
	var courseMap1 = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
		"java": "java开发工程师",
	}
	//取值方法
	fmt.Println(courseMap1["java"])
	//放值
	courseMap1["mysql"] = "mysql原理"
	//定义了一个map结构体，但是没有初始化这个map，这时候用这个map的时候，会报错
	//var courseMap2 map[string]string //这样定义是一个Nil类型的，如果直接放值的话，会报错，map类型，想要设置值，必须要先初始化
	//go语言中，有一个空的类型，叫nil，map类型想要设置值，必须先初始化，如果初始化里面没有值，可以放空值
	//courseMap2["mysql"] = "mysql的原理"
	//fmt.Println(courseMap2)

	//map的第二种初始化方式，map空初始化,map不能定义了之后，不进行初始化
	var courseMap3 = map[string]string{}
	courseMap3["姓名"] = "张三"
	fmt.Println(courseMap3)

	//map的第三种初始化方法：make函数，这种初始化的方式更为常用
	var courseMap4 = make(map[string]string, 3) //make是内置函数，主要用于初始化slice、map、channel
	courseMap4["java"] = "java学习"
	fmt.Println(courseMap4)
	//map必须初始化才能使用，map的初始化方式主要有2种：1.map[string]string{} 2.make(map[string]string)
	//但是slice可以不初始化，
	var m []string
	if m == nil {
		fmt.Println("yes")
	}
	m = append(m, "张三")
	fmt.Println(m)
	var coursesMape5 = make(map[string]string, 5)
	coursesMape5["姓名"] = "张三"
	coursesMape5["年龄"] = "18"
	coursesMape5["性别"] = "男"
	coursesMape5["身高"] = "174cm"
	coursesMape5["体重"] = "75kg"
	//第一种遍历方式
	for key, value := range coursesMape5 {
		fmt.Println(key, value)
	}
	//第二种遍历方式,这种遍历方式，用的是key
	for value := range coursesMape5 {
		fmt.Println(value)
	}
	//map是无序的，而且不保证每次打印都是相同的顺序
	d1 := courseMap1["java"] //单参数返回
	fmt.Println(d1)
	d2, ok := courseMap1["goes"]

	if !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("find...", d2)
	}

	//另一种写法
	if d, ok := courseMap1["go"]; !ok {
		fmt.Println("not in...")
	} else {
		fmt.Println("find...", d)
	}

	//删除一个元素
	delete(courseMap1, "go")
	fmt.Println(courseMap1)
	//删除一个不存在的元素，也不会报错
	delete(courseMap1, "rpc")
	//map不是线程安全的，当有两个线程goruntinue，对map操作的话，会报错，当有多个线程对一个map操作的时候，要用sunc.map包来操作

}
