package main

import "fmt"

//Definimos el struct Camisetas
type Camisetas struct {
	ID     int
	Color1 string
	Color2 string
	Team   string
}

// ⭐ IV - Crear una función que actúe como método constructor.
//La función regresa con & para que regrese la referencia y no una copia, de la struct.
func NewCamiseta(ID int, Color1 string, Color2 string, Team string) *Camisetas {
	return &Camisetas{
		ID:     ID,
		Color1: Color1,
		Color2: Color2,
		Team:   Team,
	}
}

func main() {

	//------FORMAS PARA INSTANCIAR UN OBJETO EN GO--------
	// I - con zero values
	c := Camisetas{}
	fmt.Printf("%v\n", c) //{0   }

	// II - Asignarle valores a las propiedades.
	c2 := Camisetas{
		ID:     1,
		Color1: "Negro",
		Color2: "Dorado",
		Team:   "Balanza",
	}
	fmt.Printf("%v\n", c2) //{1 Negro Dorado Balanza}

	// III - Usando la palabra reservada new.
	c3 := new(Camisetas)    //Regresa una referencia a la instancia creada, para acceder al valor usamos *e3.
	fmt.Printf("%v\n", *c3) //{0   }
	c3.ID = 1
	c3.Team = "Sub-21"
	fmt.Printf("%v\n", *c3) //{1   Sub-21}

	// ⭐ IV - Crear una función que actúe como método constructor.
	c4 := NewCamiseta(12, "Amarillo", "Azul", "Boca")
	fmt.Printf("%v\n", *c4) //{12 Amarillo Azul Boca}
}
