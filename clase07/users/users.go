package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

// Embebido
type Employee struct {
	User
	Active bool
}

// Creando Interface(comportamiento de un struct)
type Cashier interface {
	//Metodos
	CalcTotal(item ...float32) float32
	deactivate()
	reactivate()
}

type Admin interface {
	DesactivateEmployee(c Cashier)
	ReactivateEmployee(c Cashier)
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {

	if !e.Active {
		fmt.Println("Empleado desactivado")
		return 0
	}

	var sum float32

	for _, i := range item {

		sum += i
	}
	return sum * 1.15
}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) DesactivateEmployee(c Cashier) {
	c.deactivate()
}

func (e *Employee) reactivate() {
	e.Active = true
}

func (e *Employee) ReactivateEmployee(c Cashier) {
	c.reactivate()
}
