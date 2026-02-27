package main                        //声明此文件属于main包
import "fmt"                        //标准库包
import _ "github.com/gin-gonic/gin" //第三方包
//import myalias "mypackage"          //别名导入
import _ "fmt"          //点导入(可以直接使用包内函数)
import _ "database/sql" //匿名导入，仅执行init函数

/*
Package是go语言的基本组织单元，用于封装和复用代码，每个go文件都属于一个包，包名定义在文件的第一行

包的类型：

	1.可执行包Executable Package，包名必须是main，必须包含func main(),编译后生成可执行文件
	2.库包。包名可以是任何有效标识符，提供可以被其他包导入的功能，不会生成可以执行的文件

导出规则
	大写字母开头的标识符为公开exported
	小写字母开头的标识符为私有 unexported
*/

var PublicVar int = 10  //可以被外部包访问
var privateVar int = 20 //仅包内可以访问

func PublicFunc()  {} //公开函数
func privateFunc() {} //私有函数

/*
	包的初始化
		每个包都可以顶一个init()函数，按导入顺序自动执行，一个包可以有多个init()函数，按照定义顺序执行

*/

func main() {
	fmt.Println("hello world")
	Greet("张三")
}

func Greet(name string) {
	fmt.Printf("Hello,%s!\r\n", name)
}
