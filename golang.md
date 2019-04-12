<a href="#1">GOROOT GOPATH GOBIN 目录</a><br />
<a href="#2">get GitHub中的包</a><br />
<a href="#3">golang.org/x/中的包无法下载问题</a><br />
<a href="#4">exec: "gcc": executable file not found in %PATH%报错</a><br />
<a href="#5">rsrc（walk GUI依赖包）</a><br />
<a href="#6">代码包导入</a><br />
<a href="#7">代码格式化（go fmt 和 gofmt）</a>

### <a name="1">GOROOT GOPATH GOBIN 目录</a>

使用 `go env` 命令查看相关系统变量

### <a name="2">get GitHub中的包</a>

1. 系统安装了git工具
2. go get github.com/andlabs/ui   //url为包地址
3. 也可以clone文件夹到 `%GOROOT%或%GOPATH%\src\github.com\username\` 路径

### <a name="3">golang.org/x/中的包无法下载问题</a>

因为Google被墙，golang.org的相关包无法get，要曲线救国，以sys包为例：

```
git clone https://github.com/golang/sys.git
```

放置路径：`%GOROOT%或%GOPATH%\src\golang.org\x\`

- *github.com/golang 上面还有很多镜像可供墙内git*

### <a name="4">exec: "gcc": executable file not found in %PATH%报错</a>

1. 下载mingw-w64压缩包（x86_64-posix-seh）
2. 解压，把mingw-w64文件夹放在目录中（如：C:\）
3. 把路径（C:\mingw64\bin）添加到系统变量 `Path` 中

### <a name="5">rsrc（walk GUI依赖包）</a>

**安装**

```
go get github.com/akavel/rsrc
```

**编译**

```
rsrc -manifest xxxx.manifest -o rsrc.syso
```
会生成 `rsrc.syso` 文件

### <a name="6">代码包导入</a>

```
1. import str "strings"   //导入后起别名

使用：str.HasPrefix("abc", "a")

2.
import (             //导入多个代码包
  "flag"
  "fmt"
  "strings"
)

使用：strings.HasPrefix("abc", "a")

3. import . "strings"    //本地化导入

使用：HasPrefix("abc", "a")

4. import _ "strings"   //仅仅初始化，不调用程序的实体
```

### <a name="7">代码格式化（go fmt 和 gofmt）</a>

gofmt是一个独立的cli程序，go fmt命令是gofmt的简单封装
