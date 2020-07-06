package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // [April May June]
	fmt.Println(summer) // [June July August]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	fmt.Println(summer[:20]) // panic: out of range

	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // [June July August September October]
}

// 切片旋转
func TestSlice2(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 切片间比较
func TestSlice3(t *testing.T) {
	var s1 = []string{"0", "1", "2", "3"}
	var s2 = []string{"1", "2", "3"}
	var s3 = []string{"1", "2", "3"}

	fmt.Printf("%t\n", equal(s1, s2))
	fmt.Printf("%t\n", equal(s1, s3))
	fmt.Printf("%t\n", equal(s2, s3))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// 生成切片的nil值
func TestSlice4(t *testing.T) {
	var s []int
	fmt.Printf("len(s) == %d, s == %v, %t\n", len(s), s, s == nil)
	s = nil
	fmt.Printf("len(s) == %d, s == %v, %t\n", len(s), s, s == nil)
	s = []int(nil)
	fmt.Printf("len(s) == %d, s == %v, %t\n", len(s), s, s == nil)
	s = []int{}
	fmt.Printf("len(s) == %d, s == %v, %t\n", len(s), s, s == nil)
}

// 切片的append函数
func TestSlice5(t *testing.T) {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}

// 测试cap增长
func TestSlice6(t *testing.T) {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

// 切片的append函数
func TestSlice7(t *testing.T) {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}

// 优化版的appendInt函数
func TestSlice8(t *testing.T) {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendIntPlus(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendIntPlus(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// ... expand z to at least zlen...
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	copy(z[len(x):], y)
	return z
}

// 删除切片元素
func TestSlice9(t *testing.T) {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
	for k, v := range s {
		fmt.Println(k, v)
	}

	s2 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s2, 2)) // "[5 6 9 8]"
	for k, v := range s2 {
		fmt.Println(k, v)
	}
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
