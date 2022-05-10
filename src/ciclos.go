package main

import "fmt"

func main() {
	// for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")
	// for while
	count := 0
	for count <= 10 {
		fmt.Println(count)
		count++
	}

	fmt.Println("\n")
	// for forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		break
	}

	fmt.Println("\n")
	for i := 15; i >= 0; i-- {
		fmt.Println(i)
	}
}
