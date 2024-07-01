package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}
type FullTimeEmployee struct {
	Person
	Employee
	Salary int
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM PERSON WHERE DNI = dni
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM EMPLOYEE WHERE ID = id
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee = FullTimeEmployee{
		Person:   p,
		Employee: e,
		Salary:   10000,
	}

	return ftEmployee, nil // SELECT * FROM FULL_TIME_EMPLOYEE WHERE ID = id AND DNI = dni, Person{}, nil
}
