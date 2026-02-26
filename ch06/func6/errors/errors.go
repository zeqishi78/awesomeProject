package main

import (
	"fmt"
)

/*
go语言出错处理的理念，一个函数可能出错，java是用try catch去包住这个函数
开发函数的人，需要有一个返回值，去告诉调用者是否返回成功
要求go设计者，必须要处理这个error，所以代码中会大量出现if err!=nil，
go设计者认为，必须要处理这个error，这个其他语言叫做防御性编程，代码的健壮性很好，但是
*/
func main() {
	_, err := A()
	if err != nil {
		fmt.Println(err)
	}
	//B()

	var names map[string]string
	names["go"] = "go工程师"
	//为了解决问题，go给我们提供了recover函数，这个函数可以帮助我们捕获到panic
}

func A() (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover if A:", r)
		}
	}()
	return 1, nil
}

// panic是一个内置函数，这个函数会导致你的程序退出
// 在go语言中，不推荐随便使用panic,一般在哪里用到？
// 在我们的服务启动过程中，我的服务想要启动，必须有一些依赖服务必须准备好，比如日志文件存在、Mysql能链接通
// 这时候服务才能启动。如果这些服务启动检查中，发现这些任何一个都不满足，就调用panic，主动调用
// 但是你的服务一旦启动了，这时候你的某行代码中，不小心写了panic，那么不好意思你的程序就挂了，这是重大事故
// 但是架不住有些地方的代码写得不小心，会导致不被动触发panic
func B() {
	panic("this is an panic")
	fmt.Println("this is a func")

}
