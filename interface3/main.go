package main

import "fmt"

// Employee is an interface for printing employee details
type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("Diksha")
	fmt.Println("Employee Salary:", e1.PrintSalary(200000, 5))
}
