//修改echo程序，输出参数的索引和值，每行一个
package main

import (
	"os"
	"fmt"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
