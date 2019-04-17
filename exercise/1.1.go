//修改echo程序输出os.Args[0]，即命令的名字
package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Args[0])
}
