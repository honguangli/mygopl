// 练习1.3：做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
package test

import (
	"fmt"
	"strings"
	"testing"
)

var args []string

func BenchmarkEchoInefficient(b *testing.B) {
	for i := 0; i < 1000; i++ {
		args = append(args, fmt.Sprintf("%04d", i))
	}
	for i := 0; i < b.N; i++ {
		echoInefficient(args)
	}
}

func BenchmarkEchoEfficient(b *testing.B) {
	for i := 0; i < 1000; i++ {
		args = append(args, fmt.Sprintf("%04d", i))
	}
	for i := 0; i < b.N; i++ {
		echoEfficient(args)
	}
}

func echoInefficient(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	//fmt.Println(s)
}

func echoEfficient(args []string) {
	strings.Join(args, " ")
	//fmt.Println(strings.Join(args, " "))
}
