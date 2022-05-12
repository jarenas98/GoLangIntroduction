package main

import (
	structsPunteros "curso_golang_platzi/src/structsPunteros"
	"fmt"
)

func main() {

	var myPc structsPunteros.Pc = structsPunteros.Pc{Ram: 16, Disk: 512, Brand: "lenovo"}

	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)
	myPc.Ping()
}
