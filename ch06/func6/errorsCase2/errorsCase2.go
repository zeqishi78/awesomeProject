package main

import (
	"errors"
	"fmt"
	"os"
)

/*
1.返回一个错误，并将底层错误包裹在外层错误消息中，形成带上下文的错误链
2.外层信息：读取配置文件失败；底层原因：文件不存在
3.errors.Is检查是否包含某个具体根因
4.errors.As提取特定错误类型
5.
*/
func readConfig() error {
	//errors.New("文件不存在")创建一个基础错误作为根因
	//return fmt.Errorf("读取配置文件失败:%w",err)，使用%w包裹根因，形成可以解包的错误链
	//%w只能出现一次，用于错误包装：格式化后返回的类型仍为error
	return fmt.Errorf("读取配置文件失败:%w", errors.New("文件不存在"))
}

func main() {
	err := readConfig()
	if err != nil {
		fmt.Println("原始错误:", err)

		//解包错误
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			fmt.Println("解包后:", unwrapped)
		}

		//检查错误类型
		if errors.Is(err, os.ErrExist) {
			fmt.Println("文件确实不存在")
		}
	}

}
