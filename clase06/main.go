package main

import "fmt"

//Definimos el struct Camisetas
type Camisetas struct {
	ID     int
	Color1 string
	Color2 string
	Team   string
}

type parcheTitular struct {
	parche string
}

type CamisetasTitulares struct {
	Camisetas
	parcheTitular
}

func (c CamisetasTitulares) String() {
	fmt.Printf("\nid: %d, Color1: %s, Color2: %s, Team: %s, Parche: %s ", c.ID, c.Color1, c.Color2, c.Team, c.parche)
}

func main() {
	cTitular := CamisetasTitulares{}
	cTitular.ID = 45
	cTitular.Color1 = "Violeta"
	cTitular.Color2 = "Negro"
	cTitular.Team = "PujatoFC"
	cTitular.parche = "parche_conmebol_2024"
	cTitular.String()
}
