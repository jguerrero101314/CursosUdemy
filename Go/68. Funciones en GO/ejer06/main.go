package main

import "fmt"

func main() {
	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(2, 23, 45))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5, 46, 3, 4, 5, 6, 7, 20))
	// fmt.Println(uno(5))

	// numero, estado := dos(2)
	// fmt.Println(numero)
	// fmt.Println(estado)

}
func uno(numero int) (z int) {
	z = numero * 2
	return
}
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false

	} else {
		return 10, true
	}
}
func calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total

}
