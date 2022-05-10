package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Es palindroma")
	} else {
		fmt.Println("No es palindroma")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// ama
	isPalindromo("ama")
	//amor a roma
	isPalindromo("Amor a roma")
	isPalindromo("odnbi")
}
