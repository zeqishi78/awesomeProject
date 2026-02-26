# Go 学习笔记（基于 ch01–ch05 代码）

本学习文档按章节梳理知识点，并配套示例与源码位置，帮助你快速掌握 Go 的核心语法与常用标准库。

## ch01 变量、常量与作用域
- 变量定义与零值
  - 先定义后使用，类型确定后不可更改；未使用会报错
  - 基本类型的零值：int 为 0，string 为 ""，bool 为 false
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch01/main.go#L15-L27)
- 多变量定义与短变量声明
  - 同时定义多个变量；函数体内可用 `:=` 推断类型并赋值
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch01/main.go#L28-L39)
- 作用域与匿名变量
  - 代码块内定义的变量仅在块内有效；匿名变量 `_` 用于丢弃返回值
  - 源码参考：[iota/main.go](file:///d:/study/go/code/awesomeProject/ch01/iota/main.go#L11-L24)
- 常量与分组
  - 常量不可修改，推荐大写；分组可一次定义多个常量
  - 源码参考：[const/main.go](file:///d:/study/go/code/awesomeProject/ch01/const/main.go#L6-L23)
- 特殊常量 iota
  - 在 const 分组中按行递增；每遇到新的 const 分组重置为 0
  - 源码参考：[iota/main.go](file:///d:/study/go/code/awesomeProject/ch01/iota/main.go#L26-L47)

## ch02 基本类型与类型转换
- 数值类型与字符
  - byte 等价于 uint8，适合存放 ASCII 字符；rune 等价于 int32，适合 Unicode 字符
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch02/main.go#L28-L38)
- 字符串
  - 字符串不可变；与字符类型搭配输出
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch02/main.go#L39-L43)
- 强制类型转换
  - 显式转换数值类型；自定义类型间也需显式转换
  - 源码参考：[convert/main.go](file:///d:/study/go/code/awesomeProject/ch02/convert/main.go#L9-L25)
- 字符串与数值的互转
  - Atoi/Itoa、ParseX/FormatX 系列覆盖 int/float/bool 等
  - 源码参考：[convert/main.go](file:///d:/study/go/code/awesomeProject/ch02/convert/main.go#L26-L73)
- 运算符
  - 算术、逻辑、位运算与取地址
  - 源码参考：[expr/main.go](file:///d:/study/go/code/awesomeProject/ch02/expr/main.go#L6-L28)

## ch03 字符串与格式化输出
- 字符串与 Unicode
  - `len(s)` 返回字节数；`utf8.RuneCountInString(s)` 返回字符数（rune）
  - 使用 `[]rune` 正确按字符切分与遍历
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch03/main.go#L10-L13)、[compare/main.go](file:///d:/study/go/code/awesomeProject/ch03/compare/main.go#L23-L28)
- 转义与原始字符串
  - 双引号内支持转义；反引号为原始字符串，不处理转义
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch03/main.go#L15-L22)
- 输出与格式化
  - Print/Println/Printf/Sprintf 的差异；`%v` 输出值、`%T` 输出类型
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch03/main.go#L24-L50)
- 高性能拼接
  - `strings.Builder` 进行多次拼接更高效
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch03/main.go#L51-L61)
- 字符串常用操作
  - 比较（按字典序）、包含、长度、计数、分割、前后缀、索引、替换、大小写、修剪
  - 源码参考：[compare/main.go](file:///d:/study/go/code/awesomeProject/ch03/compare/main.go#L10-L54)

## ch04 流程控制
- if-else 分支
  - 基于布尔表达式分支；可嵌套与多个分支组合
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch04/main.go#L10-L36)
- for 循环
  - 经典三段式；省略形成 while 风格；双层循环示例
  - 源码参考：[for/for.go](file:///d:/study/go/code/awesomeProject/ch04/for/for.go#L9-L35)
- for range 遍历
  - 遍历字符串时 value 为字符拷贝（rune），key 为索引
  - 源码参考：[for/for.go](file:///d:/study/go/code/awesomeProject/ch04/for/for.go#L42-L73)
- 循环控制
  - break 退出、continue 跳过；演示定时递增后退出
  - 源码参考：[for/for.go](file:///d:/study/go/code/awesomeProject/ch04/for/for.go#L73-L83)
- goto 与标签
  - 在错误处理或复杂跳转场景使用（实际项目慎用）
  - 源码参考：[goto/goto.go](file:///d:/study/go/code/awesomeProject/ch04/goto/goto.go#L10-L21)
- switch 分支
  - 匹配常量分支；提供 default 兜底
  - 源码参考：[switch/switch.go](file:///d:/study/go/code/awesomeProject/ch04/switch/switch.go#L18-L29)

## ch05 数组基础
- 类型与长度
  - 数组类型包括长度，例如 `[3]string` 与 `[4]string` 类型不同
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch05/main.go#L6-L17)
- 遍历
  - 支持 `for` 与 `for range` 两种方式
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch05/main.go#L19-L21)、[main.go](file:///d:/study/go/code/awesomeProject/ch05/main.go#L43-L45)
- 初始化方式
  - 逐个赋值、字面量初始化、按索引指定、使用省略号 `...` 推断长度
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch05/main.go#L23-L42)
- 相等性比较
  - 元素类型与长度相同的数组可直接比较
  - 源码参考：[main.go](file:///d:/study/go/code/awesomeProject/ch05/main.go#L46-L49)
- 说明
  - 文中提到切片与 map（集合类型）将在后续章节进一步展开

## 补充示例
以下片段用于巩固学习，可在对应章节文件中尝试添加运行。

```go
// 计算字符串字节数与字符数
import "unicode/utf8"
s := "张三 李四 王五"
println(len(s))                      // 字节数
println(utf8.RuneCountInString(s))   // 字符数
```

```go
// 使用 strings.Builder 高效拼接
import (
  "strings"
)
var b strings.Builder
b.WriteString("用户名:")
b.WriteString("张三")
println(b.String())
```

```go
// strconv 常见转换
import "strconv"
v, _ := strconv.Atoi("42")       // "42" -> 42
s := strconv.Itoa(42)            // 42   -> "42"
f, _ := strconv.ParseFloat("3.14", 64)
h := strconv.FormatInt(255, 16)  // "ff"
```

## 学习建议
- 多动手：在每章 main.go 中增删示例，加深理解
- 观察类型与编译器提示：静态类型语言的约束是学习重点
- 关注 Unicode 与字符串处理：中文环境下尤其重要
- 逐步扩展：在数组基础上继续学习切片、map 与并发

## 详细讲解

### ch01 变量与常量细节
- 静态类型与零值
  - Go 为静态强类型：变量声明时类型确定，后续不可改变；未使用的局部变量会导致编译错误
  - 零值无需显式初始化：整型为 0，浮点型为 0.0，字符串为 ""，布尔为 false，引用类型为 nil
  - 全局变量允许未使用，但建议最小化全局状态
- 声明形式与短变量声明
  - 标准声明：var x int；显式初始化：var x int = 1；类型推断：var x = 1
  - 短变量声明仅可用于函数体内：x := 1；当左边已有同名变量时，至少有一个新变量被引入且整体为短声明表达式，才合法
  - 变量遮蔽：内层同名变量会遮蔽外层变量，阅读与维护时需注意作用域
- 作用域与匿名变量
  - 函数、代码块、for、if、switch 的限定作用域；尽量缩小变量作用域以减少副作用
  - 匿名变量 `_` 可丢弃不需要的返回值，避免“未使用变量”错误；仅丢弃值，不会分配
- 常量与未类型化常量
  - 常量在编译期确定；未类型化常量更灵活，可在使用点按上下文确定类型（如 3 可用作 int 或 float64）
  - 命名：包级导出常量用驼峰或全大写（含缩写）皆可，保持一致性；分组能提高可读性
- iota 模式
  - 在同一个 const 分组中，iota 按行递增；遇到新的 const 分组重置为 0
  - 可搭配显式赋值、占位行实现跳号；与位运算组合可定义“位标记”（例如 Flags）
  - 注意：一旦中断类型或赋值，后续行仍会使用递增后的 iota 值，但表达式需保持一致性

### ch02 类型系统与转换细节
- 整数与位宽
  - int/uint 位宽与平台相关（32 或 64 位）；跨平台或外部协议交互更推荐使用明确位宽类型（int32、int64）
  - 有符号与无符号转换需显式；溢出不报错但会截断至目标类型位宽（应避免）
- 浮点与舍入
  - float32/float64 分别对应 IEEE-754 精度；使用浮点参与比较时应考虑误差
  - 浮点转整数采用截断（向零靠拢），不是四舍五入；如 int(3.9) == 3
- 字符与字符串
  - rune 是 Unicode 码点（int32），byte 是原始字节（uint8）；字符字面量使用单引号产生 rune
  - 字符串是只读字节序列；索引与切片按字节而非字符操作，处理中文需转换为 []rune 或用 range
- strconv 家族
  - 字符串转整数：Atoi 仅十进制；ParseInt(s, base, bitSize) 支持任意进制与位宽，base 可为 0 自动识别前缀（0x、0、0b）
  - 整数转字符串：Itoa 仅十进制；FormatInt(i, base) 支持 2–36 进制
  - 浮点：ParseFloat(s, 32/64) 指定位宽；FormatFloat(f, fmt, prec, bitSize) 中 fmt 常用 'f'、'e'、'E'、'g'
  - 布尔：ParseBool 支持 "1"/"t"/"T"/"TRUE"/"true"/"True" 等多种形式；FormatBool 输出 "true"/"false"
  - 可靠性：所有 Parse 系列都可能返回错误；在解析用户输入前先做 TrimSpace，必要时做范围与格式校验
- 运算符与指针
  - ++/-- 是语句而非表达式，不能用于赋值或作为子表达式；如 a = b++ 是非法
  - 位运算：& 与、| 或、^ 异或、&^ 位清除、<< 左移、>> 右移；右移对有符号数为算术右移
  - 取地址与指针：&x 得到 *T；打印地址可用 %p；指针减少拷贝但需注意逃逸与生命周期

### ch03 字符串、Unicode 与格式化
- 字节与字符计数
  - len(s) 返回字节长度；utf8.RuneCountInString(s) 返回字符数；含中文时两者不同
  - 遍历字符串用 for range 返回索引与字符（rune）；索引单位为字节
  - 索引与切片的坑：s[i] 取的是字节；对多字节字符切片可能截断产生无效 UTF-8 序列
- 转义与原始字符串
  - 常见转义：\n 换行、\r 回车、\t 制表、\" 引号、\\ 反斜杠
  - 反引号原始字符串不处理转义，适合多行、正则或模板
- fmt 家族
  - Print/Println 直接输出；Printf 按格式化输出；Sprintf 返回字符串以便后续拼接或日志
  - 常用动词：%s 字符串、%d 整数、%f 浮点、%t 布尔、%v 值、%#v Go 语法表示、%T 类型、%q 带引号字符串/字符
  - 性能提示：Printf 格式化开销较高，频繁拼接建议 Builder 或 bytes.Buffer
- 高性能拼接
  - strings.Builder 面向字符串拼接，内部维护可增长缓冲；避免频繁创建临时字符串
  - 使用注意：不要在多个 goroutine 间共享同一个 Builder；完成后调用 String() 取结果，必要时 Reset()
- strings 常用函数
  - Contains/Count：Count 按非重叠子串计数；对单字符效果直观
  - Split/Join：Split 返回切片；用 Join 进行反向拼接
  - HasPrefix/HasSuffix：前后缀判断；Index/LastIndex 返回位置，未命中返回 -1
  - Replace：n 次替换，n=-1 表示全部；ToLower/ToUpper 做 Unicode 映射；Trim/TrimSpace 去除两端字符

### ch04 流程控制与语法惯例
- if/else
  - 条件无需括号；支持在 if 语句中先声明局部变量：if x := f(); x < 0 { ... }
  - 合理拆分分支比深层嵌套更易读；尽量“早返回”减少嵌套层级
- for 全形态
  - 经典三段式、while 风格（仅条件）、无限循环（for {}）；嵌套循环注意复杂度
  - break/continue 支持标签，便于退出多层循环
- range 语义
  - 字符串：key 为字节索引，value 为字符拷贝（rune）
  - 数组/切片：key 为索引，value 为元素拷贝；需要原地修改用索引访问
  - map：遍历键值对，迭代顺序不保证；需要稳定顺序需收集并排序 key
- goto 与标签
  - 可用于集中错误处理或复杂跳转，但现代 Go 更提倡结构化控制（defer、显式返回）
- switch
  - case 支持常量/表达式；支持多个值用逗号分隔
  - 默认不“贯穿”；若需继续执行下一分支使用 fallthrough（慎用以免降低可读性）
  - 类型 switch：switch x := any.(type) { ... } 可按动态类型分支（示例可在后续章节扩展）

### ch05 数组语义与实践
- 值语义与长度
  - 数组长度是类型的一部分；函数参数传递会拷贝整个数组（可能造成性能开销）
  - 比较需要元素类型与长度完全相同；切片则不同（引用语义）
- 初始化与遍历
  - 字面量初始化更直观；按索引指定可稀疏初始化；省略号推断长度便于维护
  - 遍历时用 len(arr) 控制边界；越界访问会 panic
- 与切片的关系
  - 切片是对底层数组的视图，包含指针、长度与容量；从数组切片获取切片后，修改切片元素会影响原数组
  - 构建需增长的序列应优先使用切片；数组适合固定大小、栈上分配与性能敏感的场景

—— 完 —
## 逐行解释

### ch01/main.go
- [ch01/main.go](file:///d:/study/go/code/awesomeProject/ch01/main.go)
- L1: 定义包为 main，表示可执行程序入口
- L3: 导入 fmt 包用于格式化与打印
- L5-L13: 注释与包级变量分组声明，定义 name、age、ok 的全局变量
- L15: main 函数程序入口
- L16-L19: 注释说明静态类型特性：先定义后使用、类型固定
- L20-L21: 注释示例（被注释掉的错误写法）
- L22: 显式声明并赋值整型变量 name1
- L23: 短变量声明 age，类型推断为 int
- L24: 注释：未使用的局部变量将报错
- L25-L26: 输出 name1 与 age
- L28-L31: 多变量定义（类型各异），并打印
- L32-L36: 注释：使用前必须定义、类型一致、变量名不能冲突
- L37-L38: 显式声明字符串 address 并打印
- L40-L42: 注释：局部变量遮蔽全局变量；变量存在零值
- L43: 打印全局变量 age
- L44-L45: 声明局部整型 age2 并打印其零值 0
- L47-L50: 注释：字符串零值为 ""，未使用变量会报错
- L51-L52: 声明字符串 name3 并打印空字符串
- L54-L55: 注释：布尔零值为 false
- L56-L57: 声明布尔 ok2 并打印 false
- L59: 结束 main

### ch01/const/main.go
- [const/main.go](file:///d:/study/go/code/awesomeProject/ch01/const/main.go)
- L1: 包 main
- L3: 导入 fmt
- L5: main 入口
- L6-L10: 常量说明与定义：PI 显式类型、PI2 隐式类型、MY_NAME 字符串；常量不可修改
- L12-L16: 常量分组定义 UNKNOW/FEMALE/MALE
- L17-L22: 常量分组与隐式复制上一行类型和值：x、y、s、z
- L23-L26: 打印 x/y/s/z
- L27-L31: 注释：常量类型范围、未使用常量不强制使用、显式类型需匹配
- L32: 结束

### ch01/iota/main.go
- [iota/main.go](file:///d:/study/go/code/awesomeProject/ch01/iota/main.go)
- L1: 包 main
- L3: 导入 fmt
- L5-L7: 定义函数 a 返回 (int, bool)，用于演示匿名变量
- L9: 包级变量 name
- L10-L13: main 入口与局部变量 mainName，打印
- L14-L17: 条件判断与块作用域变量 mname，打印
- L19-L24: 匿名变量 `_` 丢弃返回值，避免未使用错误；调用 a() 并接收 ok
- L26-L36: iota 分组定义：ERR1 从 iota+1 起递增；中间插入字符串赋值 ERR3 导致 iota 计数继续；ERR4 恢复为 iota 当前值，后续递增；ERR7 显式设为 100
- L37: 打印各常量值，观察 iota 行为
- L38-L42: 注释：中断后需显式恢复；每次 const 分组 iota 重置
- L43-L46: 新分组 ERRNEW1 的 iota=0
- L47: 打印 ERRNEW1

### ch02/main.go
- [ch02/main.go](file:///d:/study/go/code/awesomeProject/ch02/main.go)
- L1: 包 main
- L3: 导入 fmt
- L5-L23: 注释与被注释的类型示例（整数/无符号/浮点），说明类型转换与精度
- L24-L27: 注释：byte 作为字符存储类型，等价 uint8
- L28-L31: 声明 byte c，赋值 'a'+5（字符算术），按 %c 打印字符；换行
- L32-L33: 声明整型 c1=97，按 %c 打印 'a'
- L35-L37: 声明 rune c2，赋值中文字符 '张'，按 %c 打印
- L39-L41: 声明字符串 name="张三"，打印
- L43: 结束

### ch02/convert/main.go
- [convert/main.go](file:///d:/study/go/code/awesomeProject/ch02/convert/main.go)
- L1: 包 main
- L3-L6: 导入 fmt 与 strconv
- L8: main 入口
- L9-L12: 整数显式转换：int8 -> uint8
- L13-L15: 浮点到整数转换：float32 -> int32（截断）
- L17-L18: int8 转 float64
- L20: 自定义类型 IT 基于 int
- L21-L24: IT 类型变量 e、d；将基础类型显式转换为 IT；打印
- L26-L33: 字符串转数字：Atoi；当输入非数字 "a" 返回错误，打印 "conver error"，再打印零值 0
- L34-L36: 数字转字符串：Itoa
- L38-L43: ParseFloat，位宽 64，错误处理与打印
- L44-L48: ParseInt 解析十进制负数 -42
- L50-L54: ParseInt 解析八进制字符串 "12"（等于十进制 10）
- L55-L61: ParseBool 解析 "true"；错误处理与打印
- L63-L66: FormatBool 将 true 格式化为 "true"
- L68-L69: FormatFloat 以指数形式 'E' 格式化，精度 -1，位宽 64
- L71-L72: FormatInt 将 -42 按 16 进制格式化（输出 -2a），打印
- L73: 结束

### ch02/expr/main.go
- [expr/main.go](file:///d:/study/go/code/awesomeProject/ch02/expr/main.go)
- L1: 包 main
- L3: 导入 fmt
- L5: main 入口
- L6-L9: 算术运算与字符串拼接，打印结果
- L12: 取模 3%2，打印 1
- L13-L14: 自增 a++，打印新值
- L16-L20: 逻辑运算演示（if 条件成立但空块）
- L22-L25: 位运算 &（与）：60 & 13 打印 12
- L26-L27: 取地址 d := &A，打印指针值
- L28: 结束

### ch03/main.go
- [ch03/main.go](file:///d:/study/go/code/awesomeProject/ch03/main.go)
- L1: 包 main
- L3-L7: 导入 fmt、strconv、strings
- L9-L13: 字符串转 []rune 以按字符计数，打印字符数
- L15-L17: 转义字符串示例，包含 \"，打印
- L18-L19: 原始字符串示例（反引号），打印
- L21-L22: 包含 \r\n 的字符串，打印时展示换行
- L24-L25: Print 与 Println 的差异演示
- L27-L30: 字符串拼接 "hello  "+uname 并打印
- L32-L41: 使用 Printf 按格式化输出；%s/%d；强调 Printf 不自动换行；Sprintf 返回字符串再打印
- L45-L50: %v 输出切片内容；%T 输出类型（[]int 与 string）
- L51-L61: strings.Builder 高性能拼接用户名、年龄、地址、手机号，最后输出

### ch03/compare/main.go
- [compare/main.go](file:///d:/study/go/code/awesomeProject/ch03/compare/main.go)
- L1: 包 main
- L3-L7: 导入 fmt、strings、unicode/utf8
- L9-L15: 字符串比较：相等与字典序
- L17-L21: 是否包含：Contains(name, "张三")
- L23-L24: 字节长度 len(name)
- L24-L25: 字符数 utf8.RuneCountInString(name)
- L26-L28: Count 子串出现次数（非重叠）
- L29-L31: Split 分割字符串为切片
- L33-L36: 判断前缀 HasPrefix
- L37-L39: 判断后缀 HasSuffix
- L41-L43: Index 查询子串位置，未命中返回 -1
- L45-L47: Replace 子串替换，-1 表示全部
- L48-L50: 大小写转换 ToLower/ToUpper
- L52-L53: Trim 去除两端空格字符
- L54: 结束

### ch04/main.go
- [ch04/main.go](file:///d:/study/go/code/awesomeProject/ch04/main.go)
- L1: 包 main
- L3: 导入 fmt
- L5-L9: 注释：if 基本结构
- L10-L13: 定义 country 与 age；进入条件判断
- L14-L17: 年龄小于 18 且国家为中国，打印“未成年”
- L18-L22: else if 与 else 分支打印
- L24-L26: 注释：实际应用中的拆分判断
- L27-L36: 拆分成三个独立 if，分别判断并打印
- L38: 结束

### ch04/for/for.go
- [for/for.go](file:///d:/study/go/code/awesomeProject/ch04/for/for.go)
- L1: 包 main
- L3-L6: 导入 fmt 与 time
- L8-L14: 累加 1..100 的和，打印
- L15-L21: 注释与死循环示例（被注释）
- L23-L28: while 风格循环：在 i<3 条件下每秒打印并递增
- L29-L35: 双层循环打印 9×9 乘法口诀
- L37-L41: 注释：for range 语法说明
- L42-L46: 遍历字符串，打印索引与 rune 值
- L47-L50: 使用匿名变量忽略索引，仅打印字符
- L52-L58: 注释：不同类型下的 range 语义
- L60-L64: 通过索引访问字符（字节），打印索引与字节转换的字符
- L66-L72: 转为 []rune 后按字符遍历，避免中文截断
- L73-L83: 无限循环，每秒递增 round，超过 20 时 break 退出

### ch04/goto/goto.go
- [goto/goto.go](file:///d:/study/go/code/awesomeProject/ch04/goto/goto.go)
- L1: 包 main
- L3: 导入 fmt
- L5-L8: 注释：goto 使用场景与注意
- L9-L18: 双层循环，遇到 j==2 时跳转到标签 over；否则打印 i、j
- L19-L20: 标签 over：统一跳转位置
- L22: 结束

### ch04/switch/switch.go
- [switch/switch.go](file:///d:/study/go/code/awesomeProject/ch04/switch/switch.go)
- L1: 包 main
- L3: 导入 fmt
- L5-L17: 注释：switch 基本结构
- L18: 定义 day
- L19-L27: 按 day 进行分支匹配并打印；未匹配走 default
- L29: 结束

### ch05/main.go
- [ch05/main.go](file:///d:/study/go/code/awesomeProject/ch05/main.go)
- L1: 包 main
- L3: 导入 fmt
- L5-L7: 注释：集合类型概览（数组、切片、map、list）
- L8-L12: 声明两个不同长度的数组，打印各自类型
- L13-L16: 赋值数组元素并打印数组
- L17-L21: 使用 for range 遍历并打印元素
- L23-L27: 数组初始化方式 1：字面量赋值
- L28-L32: 初始化方式 2：简短声明与字面量
- L33-L37: 初始化方式 3：按索引稀疏赋值，其他为零值
- L38-L42: 初始化方式 4：使用省略号推断长度，并遍历打印
- L43-L45: 经典 for 遍历，按索引访问元素
- L46-L49: 比较两个数组是否相等（长度与元素均相同才为 true）
- L51: 结束

—— 完 —
