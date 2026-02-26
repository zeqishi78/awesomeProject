package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//字符串的比较
	a := "hello"
	b := "hello"
	c := "bello"
	fmt.Println(a == b)

	//字符串的大小比较
	fmt.Println(a > c)
	//是否包含
	name := "张三 李四 王五"
	contains := strings.Contains(name, "张三")
	fmt.Println(contains)
	//字符串长度
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))

	//查找字符串出现的次数
	fmt.Println(strings.Count(name, "张"))

	//分割字符串
	splitArr := strings.Split(name, " ")
	fmt.Println(splitArr)

	//字符串是否包含前缀、是否包含后缀？
	prefix := strings.HasPrefix(name, "张")
	fmt.Println(prefix)

	//判断字符串是否以 王五结尾？
	suffix := strings.HasSuffix(name, "王五")
	fmt.Println(suffix)

	//查询子串出现的位置
	index := strings.Index(name, "李四")
	fmt.Println(index)

	//子串替换
	fmt.Println(strings.Replace(name, "张三", "王八大", -1))

	//大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))

	//去掉特殊字符,trim主要用来去掉左右两边的字符
	fmt.Println(strings.Trim(name, " "))
}
