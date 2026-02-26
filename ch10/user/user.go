package user

//package是用来组织源码，是多个go源码的集合，代码复用的基础，比如fmt/os/io
//每个源文件，开始的时候，都要声明一个package，这个源码属于哪个包下面的
//python中不需要声明package

type Course struct {
	Name string
}
