package main

import (
	myPackage "curso_golang_platzi/src/myPackage"
	"fmt"
)

func main() {
	var myCar myPackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	myPackage.PrintMessage("Hola Platzi")
}
