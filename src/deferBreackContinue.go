package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue
	for i := 0; i < 10; i++ {
		// continer
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// break

		if i == 8 {
			fmt.Println("Break")
			break
		}

		fmt.Println(i)

	}

}
