package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		-----------------------GORUTINES------------------------
	*/

	//crear canal
	c := make(chan int)

	//disparar gorutine. no olvidar palabra reservada go
	go doSomething(c)
	<-c // // Recibir el valor del canal
}

// Creamos una funcion que simule una gorutine
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
