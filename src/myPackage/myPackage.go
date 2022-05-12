package myPackage

import "fmt"

// CarPublic car with public access
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessage(text string) {
	fmt.Println(text)
}
