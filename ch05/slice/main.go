package main

import "fmt"

func main() {
	//go语言中，数组长度不一样，类型都不一样，go切片本质也是使用的数组
	var courses []string
	fmt.Printf("%T\r\n", courses) //[]string，切片

	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")
	fmt.Println(courses)
	fmt.Println(courses[2])

	//切片的初始化方式有三种
	//切片初始化的第一种方式：从数组直接创建
	allCourses := [5]string{"go", "grpc", "gin", "docker", "elasticsearch"}
	//这里有4门课程，但是我只想把第二、三门课程取出来，做成一个slice
	courseSlice := allCourses[1:3] //左闭右开的区间
	fmt.Println(courseSlice)
	//取所有的值
	courses1 := allCourses[0:len(allCourses)]
	fmt.Println(courses1)

	//切片初始化的第二种方式：使用string{}
	courses2 := []string{"go", "grpc", "gin", "docker", "elasticsearch"}
	fmt.Println(courses2)
	//切片初始化的第三种方式：make函数
	allCourses3 := make([]string, 3)
	allCourses3[0] = "张三"
	allCourses3[1] = "李四"
	allCourses3[2] = "王五"
	fmt.Println(allCourses3)

	var allCourses4 []string
	//allCourses4[0] = "c" //这里会报错，因为底层切片空间都没有
	//对于切片初始化，如果初始的时候不声明空间大小，只能使用append的方式添加
	allCourses4 = append(allCourses4, "张三")
	allCourses4 = append(allCourses4, "王五")
	allCourses4 = append(allCourses4, "李四")
	allCourses4 = append(allCourses4, "赵六")
	allCourses4 = append(allCourses4, "胡七")
	allCourses4 = append(allCourses4, "朱八")
	fmt.Println(allCourses4)

	//访问切片的元素，访问单个，访问多个
	fmt.Println(allCourses4[1]) //访问单个
	//访问切片中的多个元素，[start:end]，如果只有start，没有end，就表示从start开始到结尾的所有数据；如果只有end，那么就是取从开头到end的所有数据
	fmt.Println(allCourses4[0:2])
	//没有end，就表示从start开始到结尾的所有数据
	fmt.Println(allCourses4[2:])
	//如果只有end，那么就是取从开头到end的所有数据
	fmt.Println(allCourses4[:3])
	//如果没有start，也没有end，就是表示复制了一份数据
	fmt.Println(allCourses4[:])

	//往一个切片中，添加另一个切片的元素
	//第一种方式
	courses3 := []string{"go", "grpc"}
	courses3 = append(courses3, "gin", "mysql", "es", "java")
	courses4 := []string{"张三", "李四", "王五"}
	for _, value := range courses3 {
		courses4 = append(courses4, value)
	}
	fmt.Println(courses4)
	//第二种添加方式
	courses4 = append(courses4, courses3...)
	fmt.Println(courses4)

	//如何删除slice中的元素，比较麻烦
	courses5 := []string{"go", "grpc", "mysql", "es", "gin", "docker", "k8s"}
	fmt.Println(courses5)
	//拼接
	myslice := append(courses5[:2], courses5[3:]...) //courses5[3:]...相当于把courses5中的元素打散，
	fmt.Println(myslice)
	courses5 = courses5[:3]
	fmt.Println(courses5)

	//复制slice
	coursesCopy := courses5
	coursesCopy1 := courses5[:] //这种方式，如果切片中元素有改动，拷贝的数据会有影响
	fmt.Println(coursesCopy, coursesCopy1)

	//将右边的切片，拷贝到左边来
	var courseSliceCopy = make([]string, len(courses5))
	copy(courseSliceCopy, courses5) //这种方式，如果切片中元素有改动，拷贝的数据不会有影响
	fmt.Println(courseSliceCopy)
	fmt.Println("=================================")
	courses5[0] = "java"
	fmt.Println(coursesCopy) //
	fmt.Println(courseSliceCopy)

}
