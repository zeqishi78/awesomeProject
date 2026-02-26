package main

import "fmt"

func a() (int, bool) {
	return 0, false
}

var name = "bobby" //全局变量
func main() {
	//变量的作用阈，代码块中定义的变量，代码块外面是拿不到变量的
	var mainName = "main"
	fmt.Println(mainName)
	if name == "bobby" {
		mname := "imooc"
		fmt.Println(mname)
	}

	//匿名变量，不使用
	var _ int
	_, ok := a()
	if ok {
		//打印
	}

	//iota：特殊常量，常量领域的关键字，可以认为是一个可以被编译器修改的常量

	const (
		ERR1 = iota + 1
		ERR2
		ERR3 = "ha" //iota内部仍然会增加计数器
		ERR25
		ERR4 = iota
		ERR5
		ERR6
		ERR7 = 100
	)
	fmt.Println(ERR1, ERR2, ERR3, ERR25, ERR4, ERR5, ERR6)
	/*
		如果中断了IOTA，那么久必须显式恢复，后续会自动递增
		每次出现const关键字的时候，iota初始化为0

	*/
	const (
		ERRNEW1 = iota
	)
	fmt.Println(ERRNEW1)
}
