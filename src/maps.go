package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// looking for
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
