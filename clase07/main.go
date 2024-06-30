package main

import (
	"clase07/users"
	"fmt"
)

func main() {
	//creando usuario
	var frank users.Cashier = users.NewEmployee("Frank")
	var robert users.Admin = users.NewEmployee("Robert")

	total := frank.CalcTotal(1, 2, 3)
	fmt.Println(total) //6.8999996

	robert.DesactivateEmployee(frank)
	fmt.Println(frank.CalcTotal(1, 2, 3)) //Empleado desactivado 0

	robert.ReactivateEmployee(frank)
	fmt.Println(frank.CalcTotal(1, 20, 3)) //27.599998
}
