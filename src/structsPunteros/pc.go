package structsPunteros

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (anyPc Pc) Ping() {
	fmt.Println(anyPc.Brand, "pong!")
}

func (anyPc *Pc) DuplicateRam() {
	anyPc.Ram = anyPc.Ram * 2
}

func (anyPc Pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", anyPc.Ram, anyPc.Disk, anyPc.Brand)
}
