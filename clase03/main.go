package main

import "fmt"

func main() {
	/*
		-----------APUNTADORES/PUNTEROS-----------
	*/
	g := 25
	fmt.Println(g)
	h := &g        // h es un apuntador de g. El valor de h es la direccion de memoria de g
	fmt.Println(h) //0xc00000a0b8
	//El operador * indica a que valor apunta esa direccion de memoria
	fmt.Println(*h) //25
}
