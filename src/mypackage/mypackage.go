package mypackage

import (
	"fmt"
)

// CarPublic Car de acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// PrintMessage
func PrintMessage(text string) {
	fmt.Println(text)
}
