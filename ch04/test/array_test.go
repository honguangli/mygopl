package test

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

// 默认初始化为元素类型对应的零值
func TestArray(t *testing.T) {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

// 使用数组字面值语法初始化数组
func TestArray2(t *testing.T) {
	/*
		var q  = [3]int{1, 2, 3}
		var r [3]int = [3]int{1, 2}
		fmt.Println(r[2]) // "0"
	*/

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	/*
		q := [3]int{1, 2, 3}
		q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
	*/
}

// 使用数组字面值语法初始化数组
// 指定索引的对应值
func TestArray3(t *testing.T) {
	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) //  "3 ￥"

	r := [...]int{99: -1}
	fmt.Printf("%T %d\n", r, len(r))
}

// 数组间比较
func TestArray4(t *testing.T) {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	/*
		d := [3]int{1, 2}
		fmt.Println(a == d) // "compile error: cannot compare [2]int == [3]int
	*/
}

// 通过指针传递数组，并将数组清零
func TestArray5(t *testing.T) {
	a := sha256.Sum256([]byte("Test"))
	fmt.Printf("%T %[1]v\n", a)

	zero(&a)
	fmt.Printf("%T %[1]v\n", a)
}
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// 利用数组字面值初始化
func zeroPlus(ptr *[32]byte) {
	*ptr = [32]byte{}
}
