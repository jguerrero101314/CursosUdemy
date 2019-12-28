package main

import "fmt"

//vectores
// var tabla [10]int
// var tabla [10]int{5, 6, 1, 2, 3, 4, 5, 6, 7, 8}

//matrices
// var matriz [5][7]int

//slices

func main() {
	//vectores
	// tabla := [10]int{5, 6, 1, 2, 3, 4, 5, 6, 7, 8}

	// for i := 0; i < len(tabla); i++ {
	// 	fmt.Println(tabla[i])
	// }
	//matrices
	// matriz[3][5] = 1
	// fmt.Println(matriz)

	//slices
	// matriz := []int{2, 5, 4}
	// fmt.Println(matriz)

	// variante2()
	variante3()
	variante4()
}

// func variante2() {
// 	elementos := [5]int{1, 2, 3, 4, 5}
// 	porcion := elementos[3:]
// 	fmt.Println(porcion)
// }
func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capaidad %d", len(elementos), cap(elementos))

}
func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
