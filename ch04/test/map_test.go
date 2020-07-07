package test

import "testing"

// 创建map
func TestMap(t *testing.T) {
	// make函数构造
	args := make(map[string]int) // mapping from strings to ints

	// 字面值语法
	args := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	args := make(map[string]int)
	args["alice"] = 31
	args["charlie"] = 34
}
