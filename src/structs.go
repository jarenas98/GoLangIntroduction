package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//otra forma
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 1998

	fmt.Print(otherCar)
}
