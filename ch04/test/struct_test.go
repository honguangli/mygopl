package test

import (
	"fmt"
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

func TestUse(t *testing.T) {
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
