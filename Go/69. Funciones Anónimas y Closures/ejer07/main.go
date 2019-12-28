package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	//restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6, 4))

	//multiplicamos
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Printf("Multiplicamos 2 * 2 = %d \n", Calculo(2, 2))

	//dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos 12 / 3 = %d \n", Calculo(12, 3))

	operaciones()
	//closure
	tabladel := 2
	MiTabla := tabla(tabladel)

	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}
func operaciones() {
	resultado := func() int {
		var a int = 2
		var b int = 2

		return a + b
	}
	fmt.Println(resultado())
}

//closures
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
