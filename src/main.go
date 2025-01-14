package main

import (
	"fmt"
	"math"
)

//Antes tenia otro archivo pero lo perdi, este es uno hecho desde un reto de Platzi

func main() {
	//Reto: Areas de Rectangulo, Trapecio y Circulo
	//r
	br := 20
	hr := 10

	ar := br * hr
	fmt.Println("Area del rectangulo: ", ar)
	//t
	Bt := 20
	bt := 14
	ht := 5

	at := ((bt + Bt) / 2) * ht
	fmt.Println("Area del trapecio: ", at)
	//c
	rc := 10
	pi := 3.14159265359
	rc2 := math.Pow(float64(rc), 2)

	ac := pi * rc2
	fmt.Println("Area del circulo: ", ac)
}
