// 练习1.2：修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)

func main() {
	for k, arg := range os.Args {
		fmt.Println(k, arg)
	}
}
