package main

import (
	"fmt"
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T) {
	tests := map[string]struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		"GetFullTimeEmployeeById": {
			id:  1,
			dni: "12345678",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "12345678",
						Name: "Valentin Zoia",
						Age:  18,
					}, nil
				}

			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					DNI:  "12345678",
					Name: "Valentin Zoia",
					Age:  18,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
				Salary: 10000,
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for name, test := range tests {
		test.mockFunc()
		employee, err := GetFullTimeEmployeeById(test.id, test.dni)
		fmt.Println(employee)
		fmt.Println(test.expectedEmployee)
		if err != nil {
			t.Error(err)
		}
		if employee.Age != test.expectedEmployee.Age {
			t.Errorf("%s: expected %d, got %d", name, test.expectedEmployee.Age, employee.Age)
		}

		if employee.Name != test.expectedEmployee.Name {
			t.Errorf("%s: expected %s, got %s", name, test.expectedEmployee.Name, employee.Name)
		}

		if employee.Id != test.expectedEmployee.Id {
			t.Errorf("%s: expected %d, got %d", name, test.expectedEmployee.Id, employee.Id)
		}

		if employee.Position != test.expectedEmployee.Position {
			t.Errorf("%s: expected %s, got %s", name, test.expectedEmployee.Position, employee.Position)
		}

		if employee.Salary != test.expectedEmployee.Salary {
			t.Errorf("%s: expected %d, got %d", name, test.expectedEmployee.Salary, employee.Salary)
		}

		GetEmployeeById = originalGetEmployeeById
		GetPersonByDNI = originalGetPersonByDNI
	}
}
