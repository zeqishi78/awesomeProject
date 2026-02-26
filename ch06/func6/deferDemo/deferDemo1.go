package main

import "fmt"

/*
1.期望结果

	(1)main开始
	(2)example开始...
	(3)example即将执行return
	(4)example中的defer开始执行了...
	(5)结果:100

defer在return之后执行，但是在函数放回给调用者之前执行
(1)return语句先执行，确定返回值
(2)defer在return之后执行，但是还在函数内部
(3)defer执行完成，函数才真正结束返回

2.实际结果：
*/
func main() {
	fmt.Println("main开始")
	result := example()
	fmt.Println("结果:", result)
}

func example() int {
	fmt.Println("example开始...")
	defer fmt.Println("example中的defer开始执行了...")
	fmt.Println("example即将执行return")
	return 100
}
