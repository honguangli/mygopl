package test

import (
	"fmt"
	"image/gif"
	"testing"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// 结构体使用
func TestStructUse(t *testing.T) {
	var dilbert Employee
	dilbert.Salary -= 5000 //  demoted, for writing too few lines of code

	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for ... no real reason
}

func EmployeeByID(id int) *Employee {
	var e = Employee{
		ID:       id,
		Position: "Pointy-haired boss",
	}
	return &e
}

// 空结构体
func TestStructEmpty(t *testing.T) {
	seen := make(map[string]struct{}) // set of strings

	// ...
	s := "s"
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
}

// 结构体字面值语法
func TestStructLiteral(t *testing.T) {
	p := Point{1, 2}

	fmt.Printf("%v\n", p)

	var nframes = 1

	anim := gif.GIF{LoopCount: nframes}

	fmt.Printf("%v\n", anim)

	fmt.Println(Scale(Point{1, 2}, 5))

	pp := &Point{1, 2}

	/*
		pp := new(Point)
		*pp = Point{1, 2}
	*/

	fmt.Println(pp)
}

type Point struct {
	X, Y int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

// 结构体比较
func TestStructCompare(t *testing.T) {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"

	type address struct {
		hostname string
		port     int
	}

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

// 结构体成员
func TestStructMember(t *testing.T) {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	/*	var w Wheel
		w.Circle.Center.X = 8
		w.Circle.Center.Y = 8
		w.Circle.Radius = 5
		w.Spokes = 20*/

	/*	var w Wheel
		w.X = 8      // equivalent to w.Circle.Point.X = 8
		w.Y = 8      // equivalent to w.Circle.Point.Y = 8
		w.Radius = 5 // equivalent to w.Circle.Radius = 5
		w.Spokes = 20*/

	w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
}

type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

// 嵌入成员
/*type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}*/

// 嵌入匿名成员
/*type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}*/
