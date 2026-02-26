package main

import (
	"fmt"
	"os"
)

/*
panic用于处理程序无法继续执行的严重错误
*/
func main() {
	//必须存在的文件
	file := mustOpenFile("config.json")
	defer file.Close()
	//如果文件不存在，程序会panic

}

func mustOpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("无法打开文件%s:%v", filename, err))
	}
	return file
}
