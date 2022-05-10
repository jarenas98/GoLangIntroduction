package main

import (
	"fmt"
)

func main() {

	fmt.Println("Declaracion de constantes")
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.145

	fmt.Println("pi : ", pi)
	fmt.Println("pi2 : ", pi2)

	fmt.Println("Declaracion de variables enteras")
	// Declaracion variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base : ", base)
	fmt.Println("altura : ", altura)
	fmt.Println("area : ", area)

	fmt.Println("Zero Values")
	//Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	fmt.Println("Area de un cuadrado")
	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado = ", areaCuadrado)
}
