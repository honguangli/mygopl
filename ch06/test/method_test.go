package test

import (
	"fmt"
	"image/color"
	"mygopl/ch06/coloredpoint"
	"mygopl/ch06/geometry"
	"sync"
	"testing"
)

func TestMethod(t *testing.T) {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))           // "5", method call
}

func TestMethod2(t *testing.T) {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
	fmt.Println(perim.Distance())             // "12", method of geometry.Path
}

func TestMethodPointer(t *testing.T) {
	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	p := geometry.Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	p = geometry.Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p) // "{2, 4}"
}

func TestMethodNil(t *testing.T) {

}

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func TestMethod3(t *testing.T) {
	var cp coloredpoint.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredpoint.ColoredPoint{geometry.Point{1, 1}, red}
	var q = coloredpoint.ColoredPoint{geometry.Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
}

func TestMethodCache(t *testing.T) {
	fmt.Println(Lookup("hhh"))
	mapping["hhh"] = "h1h1h1"
	fmt.Println(Lookup("hhh"))

	fmt.Println(Lookup2("hhh"))
	cache.mapping["hhh"] = "h1h1h1"
	fmt.Println(Lookup2("hhh"))
}

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
