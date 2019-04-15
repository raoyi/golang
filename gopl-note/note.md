# 《Go程序设计语言》笔记

本书源码仓库：https://github.com/adonovan/gopl.io

## 第一章

### 1.1 hello, world

1. `go run helloworld.go` 直接输出结果

2. `go build helloworld.go` 生成可随时执行的二进制文件

3. `go get gopl.io/ch1/helloworld` 获取第三方源代码

*go工具见10.7节*

4. 必须精确导入需要的包，否则出现编译错误

5. 顺序：`package`，`import`，`函数/变量/常量声明`

6. Go语言不需要分号结尾。（多个语言处于同一行除外）

7. `{` 符号必须和关键字在同一行，不能独自成行

8. `gofmt` / `go fmt子命令`（自动gofmat的IDE，如带有Go-plus插件的Atom）

9. `goimports` 工具管理导入声明的插入和移除。获取路径：

  `go get golang.org/x/tools/cmd/goimports`

### 1.2 命令行参数

1. os包  (命令行参数：os.Args  字符串slice)

2. 所有索引使用半开区间（包含第一个索引，不包含最后一个索引）

  `os.Args[0]`：命令本身

  `os.Args[1]`：第一个参数，以此类推

  `len(os.Args)`：os.Args数组的长度

  *`s[m:n]` 表示一个从第`m`个到第`n-1`个元素的slice*

  *如果 `m` 或 `n` 缺失，默认分别为 `0` 或 `len(s)`*

  `os.Args[1:len(Args)]`：命令的所有参数

  所以，以上命令可简写成：`os.Args[1:]`

3. 变量声明：

```
var <变量名，用逗号分隔多个变量名> <变量类型>
```

4. 若 s, sep, os.Args[i] 均为字符串

```
s += sep + os.Args[i] //等价于
s = s + sep + os.Args[i]
//将srp和os.Args[i]追加到旧的s上，并重新赋值给s
//+= 为 赋值操作符，类似的还有 *=
```

5. 短变量声明（`:=`）声明一个或多个变量，并且根据初始化的值给予合适的类型

6. `i++` 等价于 `i += 1` 等价于 `i = i + 1`

  **注意：`j = i++` 不合法，`++i` 不合法（i++是语句，不是表达式）**

7. 唯一的循环语句：`for`
