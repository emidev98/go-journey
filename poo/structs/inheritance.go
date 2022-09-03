package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	startDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	endDate string
}

type EmployeeInfo interface {
	getInfo() string
}

func (ftEmployee FullTimeEmployee) getInfo() string {
	return "Full time employee"
}

func (tEmployee TemporaryEmployee) getInfo() string {
	return "Temporary employee"
}

func printEmployeeInfo(p EmployeeInfo) {
	fmt.Println(p.getInfo())
}

func main() {
	ftEmpl := FullTimeEmployee{}
	tEmployee := TemporaryEmployee{}

	printEmployeeInfo(ftEmpl)
	printEmployeeInfo(tEmployee)
}
