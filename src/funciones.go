package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func multipleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")

	multipleArgument(1, 2, " hola")

	fmt.Println("Value = ", returnValue(3))

	value1, value2 := doubleReturn(3)
	println("value 1 = ", value1)
	println("value 2 = ", value2)

	Value1, _ := doubleReturn(3)
	println(Value1)

}
