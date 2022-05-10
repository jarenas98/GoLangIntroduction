package main

import "fmt"

func main() {
	// declaración de variables
	var helloMessages string = "Hello"
	var worldMessage string = "world"

	//Println
	fmt.Println(helloMessages, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cusos \n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cusos \n", nombre, cursos)

	// Springtf
	message := fmt.Sprintf("%s tiene más de %d cusos \n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos

	fmt.Printf("hello message: %T\n", helloMessages)
	fmt.Printf("hello message: %T\n", cursos)

}
