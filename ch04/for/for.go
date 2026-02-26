package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	//这三块可以灵活调整
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//死循环，和while(true)效果一样
	var i int
	//for {
	//	//time.Sleep(1 * time.Second)
	//	fmt.Println(i)
	//	i++
	//}

	//
	for i < 3 {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		i++
	}
	//乘法口诀
	for m := 1; m <= 9; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d * %d = %d\t", n, m, m*n)
		}
		fmt.Println()
	}

	//for循环，还有另一种用法，for range，主要是对于字符串、数组、切片、map、channel，进行ofr循环遍历
	/*
			for key,value := range{
		}
	*/
	name := "imooc go体系课"
	for index, value := range name {
		fmt.Println(index, value)
		fmt.Printf("%d %c\r\n", index, value)
	}
	//如果不想取位置信息，可以用匿名变量
	for _, value := range name {
		fmt.Printf("%c\r\n", value)
	}

	/*for range key,value
	如果是字符串的话，key就代表的是字符串的索引，value值代表的字符串值的拷贝
	字符串	字符串的索引key	字符串对应的索引的字符值拷贝(value)		如果不写key，那么返回的是索引
	数组		数组的索引		索引对应值的拷贝						如果不写key，那么返回的是索引
	切片		切片的索引		索引对应值的拷贝						如果不写key，那么返回的是map的值

	*/

	for index := range name {
		fmt.Println(index)
		fmt.Println(name[index])
		fmt.Printf("%c \r\n", name[index])
	}

	for _, value := range name {
		fmt.Printf("%c\r\n", value)
	}
	runes := []rune(name)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c/r/n", runes[i])
	}
	//for循环的退出：countinue；break
	round := 10
	for {
		time.Sleep(1 * time.Second)
		round++
		fmt.Println(round)
		if round > 20 {
			break
		}
	}

}
