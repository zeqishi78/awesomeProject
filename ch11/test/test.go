package main

//单元测试 go test
//go test命令是一个按照一定的约定和组织的测试的代码驱动程序，在包目录中，所有以_test.go结尾的为后缀的源码文件，都会被go test运行到
//我们写的以_test.go源码文件不用担心文件内容过多，因为go build命令不会将这些测试文件打包到最后的可执行的文件中
//test文件有4类，Test开头的是功能测试，以BenchMark开头的是性能测试，以example开头的是模糊测试
