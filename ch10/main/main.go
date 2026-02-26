package main

/*
	别名的好处：如果在其他目录下，存在相同的包，在引入的时候，就会报错，引入别名的话，就不会报错
*/
import (
	_ "awesomeProject/ch09/user"    //要引入，但是不用，用这种做法，主要是为了让你自动调用init
	uc10 "awesomeProject/ch10/user" //别名
	"fmt"
	_ "github.com/gin-gonic/gin"
)

// 引入包的路径
func main() {
	c := uc10.Course{
		Name: "hello java",
	}
	fmt.Println(uc10.GetCourse(c))
}
