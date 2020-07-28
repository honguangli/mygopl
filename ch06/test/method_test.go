package test

import (
	"fmt"
	"mygopl/ch06/geometry"
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
