package main

import "fmt"

//Definimos el struct Camisetas
type Camisetas struct {
	ID     int
	Color1 string
	Color2 string
	Team   string
}

// func agregarCamiseta(id int, color1, color2, team string) Camisetas{

// }

func main() {

	camiseta := Camisetas{ID: 10, Color1: "Azul", Color2: "Rojo", Team: "San lorenzo"}
	fmt.Println(camiseta, camiseta.ID, camiseta.Team)

}
