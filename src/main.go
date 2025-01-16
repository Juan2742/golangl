package main

import "fmt"

//Antes tenia otro archivo pero lo perdi, este es uno hecho desde un reto de Platzi

// Reto2: Areas en funciones
// r
/*func areaRectangulo(br, hr int) int {
	arec := br * hr
	return arec
}*/
// t
/*func areaTrapecio(Bt, bt, ht int) int {
	atra := ((Bt + bt) * ht) / 2
	return atra
}*/
// c
/*func areaCirculo(rcir2, pi float64) float64 {
	acir := rcir2 * pi
	return acir
}*/

func main() {
	//Reto1: Areas de Rectangulo, Trapecio y Circulo
	//r
	/*br := 20
	hr := 10

	ar := br * hr
	fmt.Println("Area del rectangulo: ", ar)*/
	//t
	/*	Bt := 20
		bt := 14
		ht := 5

		at := ((bt + Bt) / 2) * ht
		fmt.Println("Area del trapecio: ", at)*/
	//c
	/*	rc := 10
		pi := 3.14159265359
		rc2 := math.Pow(float64(rc), 2)

		ac := pi * rc2
		fmt.Println("Area del circulo: ", ac)*/

	//Uso de fmt
	//Printf
	/*	nombre := "Platzi"
		num := 10000

		fmt.Printf("%s tiene mas de %d cursos\n", nombre, num)*/
	//Sprintf
	/*	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, num)
		fmt.Println(message)*/
	//Tipo de dato
	/*	n := "asf"
		m := 123
		fmt.Printf("n es: %T\n", n)
		fmt.Printf("m es: %T\n", m)*/

	//Otra parte del reto2
	//r
	/*	brec := 20
		hrec := 10
		arec := areaRectangulo(brec, hrec)
		fmt.Println("Area del rectangulo: ", arec)*/
	//t
	/*	Btra := 20
		btra := 14
		htra := 5
		atra := areaTrapecio(Btra, btra, htra)
		fmt.Println("Area del trapecio: ", atra)*/
	//c
	/*	rcir := 10
		pi := 3.14159265359
		rcir2 := math.Pow(float64(rcir), 2)
		acir := areaCirculo(rcir2, pi)
		fmt.Println("Area del circulo: ", acir)*/

	//Loops: for, for while, for forever
	//for
	/*	for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
		fmt.Printf("\n")*/
	//for while
	/*	counter := 0
		for counter <= 10 {
			fmt.Println(counter)
			counter++
		}*/

	//Condicionales

	//if y and
	/*	valor1 := 1
		valor2 := 2
		if valor1 == 1 && valor2 == 2 {
			fmt.Println("Si")
		} else {
			fmt.Println("No")
		}*/
	//or
	/*	valor3 := 3
		valor4 := 5
		if valor3 == 3 || valor4 == 4 {
			fmt.Println("Si")
		} else {
			fmt.Println("No")
		}*/
	//Convertir string a Int
	/*	value, err := strconv.Atoi("23")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Value: ", value)*/

	//switch
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	//Sin condicion
	value := 60
	switch {
	case value < 0:
		fmt.Println("Es menor que 0")
	case value > 100:
		fmt.Println("Es mayor que 100")
	default:
		fmt.Println("No hay condicion")
	}

}
