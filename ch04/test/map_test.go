package test

import (
	"fmt"
	"sort"
	"testing"
)

// 创建map
func TestMapCreate(t *testing.T) {
	// make函数构造
	ages := make(map[string]int) // mapping from strings to ints

	/*
		// 字面值语法
		ages2 := map[string]int{
			"alice":   31,
			"charlie": 34,
		}
	*/

	/*
		// 等价于上述字面值语法
		ages := make(map[string]int)
		ages["alice"] = 31
		ages["charlie"] = 34
	*/

	// 访问元素
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // "32

	// delete函数删除元素
	delete(ages, "alice") // remove element ages["alice"]

	ages["bob"] = ages["bob"] + 1 // happy birthday!

	ages["bob"] += 1

	ages["bob"]++

	//_ = &ages["bob"] // compile error: cannot take address of map element
}

// 遍历map
func TestMapRange(t *testing.T) {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// 排序
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	//names := make([]string, 0, len(ages))

	/*
		var ages map[string]int
		fmt.Println(ages == nil)    // "true"
		fmt.Println(len(ages) == 0) // "true"
	*/
	ages["carol"] = 21 // panic: assignment to entry in nil map
}

// 获取map元素并报告元素是否存在
func TestMapGetElementExist(t *testing.T) {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	age, ok := ages["bob"]
	if !ok {
		/* "bob" is not a key in this map; age == 0. */
	}
	fmt.Printf("%d\t\n", age)

	if age, ok := ages["bob"]; !ok {
		/* ... */
		fmt.Printf("%d\n", age)
	}
}

// 判断两个map是否包含相同键值对
func TestMapEqual(t *testing.T) {
	var ok bool
	// True if equal is written incorrectly
	ok = mapEqual(map[string]int{"A": 0}, map[string]int{"B": 42})

	fmt.Printf("%t\n", ok)
}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// 利用辅助函数实现自定义key（非可比较的类型）比较
var m = make(map[string]int)

func TestMapKFunc(t *testing.T) {
	var slice = []string{"1", "A", "a"}
	m[k(slice)] = 1

	for k, v := range m {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
