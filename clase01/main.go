package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//----------------Declaracion Variables-----------------------------
	//Explicita
	var x int
	x = 8
	//Implicita
	y := 7

	fmt.Println(x, y) // 8 7

	//----------------Manejo de errores----------------------------
	myValue, err := strconv.ParseInt("7", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)                                            //7
		fmt.Println("The type of myValue is:", reflect.TypeOf(myValue)) //The type of myValue is: int64
	}

	//Ahora pongamos un ejemplo en el que de error
	myValue2, err2 := strconv.ParseInt("NaN", 0, 64)

	if err2 != nil {
		fmt.Printf("%v\n", err2) //strconv.ParseInt: parsing "NaN": invalid syntax
	} else {
		fmt.Println(myValue2)
		fmt.Println("The type of myValue is:", reflect.TypeOf(myValue2))
	}

	/*--------------------MAPS--------------------------------------
	los maps son estructuras de llave valor, es decir que puedo asignar
	una estructura utilizanod make(). make(map[llave]valor)
	*/
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"]) //6
	fmt.Println(m)        //map[Key:6]

	/*
		---------------------SLICE------------------------------------------
		Una estructura de datos muy similar a un array que podremos utilizar como una lista
	*/
	//Declaracion Slice de enteros
	s := []int{1, 2, 3}

	//Iterar los elementos con un for
	for index, value := range s {
		fmt.Println(index, value)
	}

	//Agregar valor nuevo al slice
	//utilizando append
	s = append(s, 4, 5)
	for index, value := range s {
		fmt.Println(index, value)
	}
}
