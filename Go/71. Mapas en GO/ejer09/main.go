package main

import "fmt"

func main() {
	// paises := make(map[string]string, 5)
	// fmt.Println(paises)

	// paises["Panamá"] = "Panamá"
	// paises["Argentina"] = "Buenos aires"

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca Junior": 30}
	fmt.Println(campeonato)

	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 55
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene  un puntaje de : %d \n", equipo, puntaje)
	}
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("el puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

}
