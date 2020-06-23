// 练习1.1：修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
package test

import (
	"fmt"
	"os"
	"testing"
)

func TestEcho1(t *testing.T) {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
