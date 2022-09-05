package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		idCard           string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:     1,
			idCard: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id: 1,
					}, nil
				}

				GetPersonByIdCard = func(idCard string) (Person, error) {
					return Person{
						name:   "John Doe",
						age:    35,
						idCard: "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					age:    35,
					idCard: "1",
					name:   "John Doe",
				},
				Employee: Employee{
					id: 1,
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByIdCard := GetPersonByIdCard

	for _, item := range table {
		item.mockFunc()

		ft, err := GetFullTimeEmployeeById(item.id, item.idCard)

		if err != nil {
			t.Errorf("ERROR: employee on retrieving employee")
		}

		if ft.age != item.expectedEmployee.age {
			t.Errorf("ERROR: expected %d age, received %d age", item.expectedEmployee.age, ft.age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByIdCard = originalGetPersonByIdCard
}
