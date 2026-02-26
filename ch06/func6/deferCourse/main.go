package main

/*
defer可以理解为java中的finnally
defer在return语句之后执行，但是函数真正返回给调用者之前执行
*/

func deferReturn() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}
//func main() {
	//连接数据库、打开文件、开始锁，无论如何，不管成功失败也好，最后都要记得关闭数据库连接、关闭文件，解锁
	//var mu sync.Mutex
	//mu.Lock()
	//defer后面的代码是放在函数return之后执行
	//defer mu.Unlock()
	//一段代码有多个defer语句
	//打开文件
	//释放文件
	//defer都是在returun函数之前去做的，
	//多个defer的时候，是栈的概念，先写的defer会先进栈，后写的defer后进栈，
	//在实际执行的时候，后进栈的defer会先弹栈，会先执行
	//defer fmt.Println("第1个defer")
	//defer fmt.Println("第2个defer")
	//fmt.Println("main逻辑...")
	//ret := deferReturn()
	//fmt.Printf("ret = %d\r\n", ret) //11
	//
	//return
//}
