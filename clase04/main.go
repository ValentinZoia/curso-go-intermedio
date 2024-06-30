package main

import "fmt"

//Definimos el struct Camisetas
type Camisetas struct {
	ID     int
	Color1 string
	Color2 string
	Team   string
}

//-------------FUNCIONES PARA SETEAR INFO----------------
func (c *Camisetas) SetId(id int) {
	c.ID = id
}

func (c *Camisetas) SetColor1(color1 string) {
	c.Color1 = color1

}

func (c *Camisetas) SetColor2(color2 string) {
	c.Color2 = color2

}

func (c *Camisetas) SetTeam(team string) {
	c.Team = team

}

//------------FUNCIONES PARA OBTENER INFO--------------
func (c *Camisetas) GetId() int {
	return c.ID
}

func (c *Camisetas) GetColor1() string {
	return c.Color1

}

func (c *Camisetas) GetColor2() string {
	return c.Color2

}

func (c *Camisetas) GetTeam() string {
	return c.Team

}

func main() {

	camiseta := Camisetas{ID: 10, Color1: "Azul", Color2: "Rojo", Team: "San lorenzo"}
	fmt.Println(camiseta, camiseta.ID, camiseta.Team) //{10 Azul Rojo San lorenzo} 10 San lorenzo

	//-----setear info------
	camiseta.SetId(33)                                //modifica la variable original
	fmt.Println(camiseta, camiseta.ID, camiseta.Team) //{33 Azul Rojo San lorenzo} 33 San lorenzo

	camiseta.SetColor1("Marron") //{33 Marron Rojo San lorenzo}
	fmt.Println(camiseta)

	camiseta.SetColor2("Blanco") //{33 Marron Blanco San lorenzo}
	fmt.Println(camiseta)

	camiseta.SetTeam("All Boys")
	fmt.Println(camiseta)

	//-------obtener info--------
	fmt.Println(camiseta.GetId())     //33
	fmt.Println(camiseta.GetColor1()) //Marron
	fmt.Println(camiseta.GetColor2()) //Blanco
	fmt.Println(camiseta.GetTeam())   //All Boys

}
