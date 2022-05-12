// package main

// import "fmt"

// type pc struct {
// 	ram   int
// 	disk  int
// 	brand string
// }

// func (anyPc pc) ping() {
// 	fmt.Println(anyPc.brand, "pong!")
// }

// func (anyPc *pc) duplicateRam() {
// 	anyPc.ram = anyPc.ram * 2
// }

// func main() {
// 	a := 5
// 	b := &a

// 	fmt.Println(b)
// 	fmt.Println(*b)

// 	*b = 100

// 	fmt.Println(a)

// 	myPc := pc{ram: 16, disk: 512, brand: "msi"}
// 	fmt.Println(myPc)

// 	myPc.ping()

// 	fmt.Println(myPc)
// 	myPc.duplicateRam()
// 	fmt.Println(myPc)
// 	myPc.duplicateRam()
// 	fmt.Println(myPc)

// }
