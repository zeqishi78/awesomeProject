package main

import (
	"container/list"
	"fmt"
)

func main() {
	var myList list.List
	myList.PushBack("go")
	myList.PushBack("grpc")
	myList.PushBack("java")
	myList.PushBack("tomcat")
	myList.PushBack("docker")
	fmt.Println(myList)

	//遍历打印值
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//反向遍历
	for i := myList.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	//数据存入
	myList1 := list.New()
	myList1.PushBack("张三")
	myList1.PushBack("李四")
	myList1.PushBack("王五")
	myList1.PushBack("赵六")
	myList1.PushBack("胡七")
	fmt.Println(myList1)
	for i := myList1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//头部放数据
	myList2 := list.New()
	myList2.PushFront("语文")
	myList2.PushFront("数学")
	myList2.PushFront("英语")
	myList2.PushFront("物理")
	myList2.PushFront("化学")
	myList2.PushFront("生物")
	//遍历打印值，正序
	for i := myList2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	i := myList2.Front()
	for ; i != nil; i = i.Next() {
		if i.Next().Value.(string) == "物理" {
			break
		}
	}
	myList2.InsertBefore("政治", i)
	for i := myList2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("====================================================================")
	//反向遍历
	for i := myList2.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
	fmt.Println("====================================================================")
	j := myList2.Front()
	for ; j != nil; j = j.Next() {
		if j.Value.(string) == "化学" {
			break
		}
	}
	myList2.Remove(j)
	//反向遍历
	for i := myList2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//4种集合类型：
	//(1)数组-不同长度的数组类型不一样；
	//(2)切片-动态数组，用起来很方便，而且性能很高，我们要尽量使用；
	//(3)map
	//(4)list
}
