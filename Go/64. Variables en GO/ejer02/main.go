package main

import (
	"fmt"
	"strconv"
)

var Numero int
var Texto string
var Status bool = true

func main() {
	var numero2, numero3, numero4 int
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Hola", false

	fmt.Println(numero2)

	fmt.Println(numero3)

	fmt.Println(numero4)

	fmt.Println(texto)
	fmt.Println(status)

	var moneda int64 = 0
	texto = strconv.Itoa(int(moneda))

	//numero3 := 3
	mostrarStatus()

}
func mostrarStatus() {
	fmt.Println(Status)
}
