package main

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

// Palindromo
/*func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es un palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}*/

// Structs
/*type car struct {
	brand string
	year  int
}*/

// Struct y Punteros
/*type pc struct {
	ram   int
	disk  int
	brand string
}
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}
func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}*/

// Stringer: Personalizar output
/*type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) Stringer() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB Disco, y es una %s", myPC.ram, myPC.disk, myPC.brand)
}*/

// Interfaces y listas de interfaces
/*type figuras2D interface {
	area() float32
}

type cuadrado struct {
	base float32
}
type rectangulo struct {
	base   float32
	altura float32
}

func (c cuadrado) area() float32 {
	return c.base * c.base
}
func (r rectangulo) area() float32 {
	return r.base * r.altura
}
func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}*/

// Gopher
/*func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(text)
}*/

// Channels
/*func say(text string, c chan string) {
	c <- text

}*/

// CH: Range, Close y Select
/*func message(text string, c chan string) {
	c <- text
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
	/*	switch modulo := 5 % 2; modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
		}*/
	//Sin condicion
	/*	value := 60
		switch {
		case value < 0:
			fmt.Println("Es menor que 0")
		case value > 100:
			fmt.Println("Es mayor que 100")
		default:
			fmt.Println("No hay condicion")
		}*/

	//Uso de defer, continue y break
	//defer
	/*	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}*/
	//Continue
	/*	i := 0
		for i = 0; i <= 9; i++ {
			fmt.Println(i)
			if i == 2 {
				fmt.Println("Hay un 2")
				continue
			}*/
	//Break
	/*		if i == 3 {
				fmt.Println("Break")
				break
			}
		}*/

	//Arrays y Slices
	//Arrays
	/*	var ar [4]int
		ar[0] = 1
		ar[1] = 2
		fmt.Println(ar, len(ar), cap(ar))*/
	//Slices
	/*	sl := []int{0, 1, 2, 3, 4, 5, 6}
		fmt.Println(sl, len(sl), cap(sl))*/
	//Metodos de slice
	/*	fmt.Println(sl[0])
		fmt.Println(sl[:3])
		fmt.Println(sl[2:4])
		fmt.Println(sl[4:])*/
	//Append
	/*	sl = append(sl, 7)
		fmt.Println(sl)
		newsl := []int{8, 9, 10}
		sl = append(sl, newsl...)
		fmt.Println(sl)*/

	//Range
	/*	slst := []string{"hola", "que", "hace"}
		for i, valor := range slst {
			fmt.Println(i, valor)
		}*/

	//Palindromo
	//	isPalindromo(strings.ToLower("Amor a roma"))

	//Maps
	/*	m := make(map[string]int)

		m["Jose"] = 14
		m["Pepito"] = 20
		fmt.Println(m)*/
	//Recorrer map
	/*	for i, v := range m {
		fmt.Println(i, v)
	}*/
	//Encontrar el valor
	/*	value, ok := m["Josep"]
		fmt.Println(value, ok)*/

	//Structs
	/*	newcar := car{brand: "Ford", year: 2020}
		fmt.Println(newcar)*/
	//otra forma
	/*	var otherCar car
		otherCar.brand = "Ferrari"
		fmt.Println(otherCar)*/

	//Struct publico: mypackage.go
	/*	var myCar pk.CarPublic
		myCar.Brand = "Ferrari"
		myCar.Year = 1990
		fmt.Println(myCar)
		pk.PrintMessage("Master")*/

	//Struct y Punteros
	/*	a := 50
		b := &a

		fmt.Println(b)
		fmt.Println(*b)

		*b = 100
		fmt.Println(a)*/
	//c
	/*	myPC := pc{ram: 16, disk: 200, brand: "msi"}
		fmt.Println(myPC)

		myPC.ping()

		fmt.Println(myPC.ram)
		myPC.duplicateRAM()

		fmt.Println(myPC.ram)
		myPC.duplicateRAM()

		fmt.Println(myPC.ram)*/

	//Reto 3:
	/*	myPC := pk.Pc{Ram: 16, Disk: 200, Brand: "msi"}
		fmt.Println(myPC)

		myPC.Ping()

		fmt.Println(myPC.Ram)
		myPC.DuplicateRAM()

		fmt.Println(myPC.Ram)
		myPC.DuplicateRAM()

		fmt.Println(myPC.Ram)*/

	//Stringer: Personalizar outputs
	/*	myPC := pc{ram: 16, disk: 200, brand: "msi"}
		fmt.Println(myPC.Stringer())*/

	//Interfaces y Listas de interfases
	/*	myCuadrado := cuadrado{base: 2}
		myRectangulo := rectangulo{base: 2, altura: 4}

		calcular(myCuadrado)
		calcular(myRectangulo)

		//Listas de varios values
		myInterface := []interface{}{"Hola", 32, 4.90}
		fmt.Println(myInterface...)*/

	//Gopher
	/*	var wg sync.WaitGroup

		fmt.Println("Hola")

		wg.Add(1)
		go say("Mundo", &wg)
		wg.Wait()

		go func(text string) {
			fmt.Println(text)
		}("Adios")

		time.Sleep(time.Second * 1)*/

	//Channnels
	/*c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)
	fmt.Println(<-c)*/

	//Channels: Range, Close,  y select
	/*c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"
	//Close: buena practica
	close(c)

	fmt.Println(len(c), cap(c))
	//Range
	for message := range c {
		fmt.Println(message)
	}
	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido desde email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido desde email2", m2)
		}
	}*/

	//Framework: Echo
	/*	e := echo.New()
		//ruta
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello world")
		})

		e.Logger.Fatal(e.Start(":1323"))*/

	//Interactuar, como pedir datos al ususario
	/*	var response string
		fmt.Println("Hola, como estas?")
		fmt.Scanln(&response)
		fmt.Println("Estas", response)*/

}
