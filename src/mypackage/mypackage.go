package mypackage

import "fmt"

// CarPublic Car de acceso publico
/*type CarPublic struct {
	Brand string
	Year  int
}

// PrintMessage
func PrintMessage(text string) {
	fmt.Println(text)
}*/

//Reto 3: PC

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "Pong")
}
func (myPC *Pc) DuplicateRAM() {
	myPC.Ram = myPC.Ram * 2
}
