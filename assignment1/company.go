package company

import (
	"fmt"
)

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d", f.ID, f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d, Hours Worked: %.1f", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(id string, emp Employee) {
	c.Employees[id] = emp
}

func (c *Company) ListEmployees() {
	for id, emp := range c.Employees {
		fmt.Printf("%s: %s\n", id, emp.GetDetails())
	}
}
