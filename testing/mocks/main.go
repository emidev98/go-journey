package main

import (
	"time"
)

type Person struct {
	name   string
	age    int
	idCard string
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

type TemporaryEmployee struct {
	Person
	Employee
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

var GetPersonByIdCard = func(idCard string) (Person, error) {
	time.Sleep(1 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(1 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, idCard string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)

	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e
	p, err := GetPersonByIdCard(idCard)

	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p

	return ftEmployee, nil
}
