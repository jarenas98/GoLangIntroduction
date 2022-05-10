package main

import "fmt"

func main() {
	x := 10
	y := 50

	// suma
	resultado := x + y
	fmt.Println("Suma = ", resultado)

	// resta
	resultado = y - x
	fmt.Println("Resta = ", resultado)

	// multiplicacion
	resultado = x * y
	fmt.Println("multiplicacion = ", resultado)

	// division
	resultado = y / x
	fmt.Println("Division = ", resultado)

	// modulo
	resultado = y % x
	fmt.Println("Modulo = ", resultado)

	// incremental
	x++
	fmt.Println("Incremental = ", x)

	// Decremental
	y--
	fmt.Println("Decremental = ", y)

	// Area de un rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10
	fmt.Println("Área rectangulo = ", baseRectangulo*alturaRectangulo)

	// Área trapecio
	base1Trapecio := 20
	base2Trapecio := 10
	alturaTrapecio := 5
	resultado = ((base1Trapecio + base2Trapecio) / 2) * alturaTrapecio
	fmt.Println("Área trapecio = ", resultado)

	// Área círculo
	var radioCirculo float64 = 10
	phi := 3.141516

	resultadoCirculo := phi * (radioCirculo * radioCirculo)
	println("Área circulo = ", resultadoCirculo)

}
