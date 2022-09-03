package main

import "fmt"

type Employee struct {
	id        int
	name      string
	vacations bool
}

func NewEmployee(id int, name string, vacations bool) *Employee {
	return &Employee{id, name, vacations}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) SetVacations(vacations bool) {
	e.vacations = vacations
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetVacations() bool {
	return e.vacations
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)

	e2 := Employee{
		id:        1,
		name:      "Hello World",
		vacations: true,
	}
	fmt.Printf("%v\n", e2)

	e3 := new(Employee)
	e3.id = 1
	e3.name = "dlroW olleH"
	fmt.Printf("%v\n", e3)

	e4 := NewEmployee(1, "World Hello", true)
	fmt.Printf("%v\n", e4)

}
