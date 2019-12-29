package main

import "fmt"

//CASO 1
/* type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}
type animal interface {
	respirar()
	comer()
	Escarnivoro() bool
}
type vegetal interface {
	ClasificacionVegetal() string
}*/

/* Genero Humano */
/*
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}
type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

func (h *hombre) respirar()    { h.respirando = true }
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string { return "Hombre" }

func (h *mujer) respirar()    { h.respirando = true }
func (h *mujer) comer()       { h.comiendo = true }
func (h *mujer) pensar()      { h.pensando = true }
func (h *mujer) sexo() string { return "Mujer" }

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func main() {
	Joel := new(hombre)
	HumanosRespirando(Joel)

	Ruth := new(mujer)
	HumanosRespirando(Ruth)
}*/

//CASO 2
/*
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}
type animal interface {
	respirar()
	comer()
	Escarnivoro() bool
}
type vegetal interface {
	ClasificacionVegetal() string
}*/

/* Genero Humano */
/*
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}
type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}

}

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func main() {
	Joel := new(hombre)
	Joel.esHombre = true
	HumanosRespirando(Joel)

	Ruth := new(mujer)
	HumanosRespirando(Ruth)
} */

//CASO 3
// type humano interface {
// 	respirar()
// 	pensar()
// 	comer()
// 	sexo() string
// }
// type animal interface {
// 	respirar()
// 	comer()
// 	Escarnivoro() bool
// }
// type vegetal interface {
// 	ClasificacionVegetal() string
// }

/* Genero Humano */

// type hombre struct {
// 	edad       int
// 	altura     float32
// 	peso       float32
// 	respirando bool
// 	pensando   bool
// 	comiendo   bool
// 	esHombre   bool
// }
// type mujer struct {
// 	hombre
// }

// func (h *hombre) respirar() { h.respirando = true }
// func (h *hombre) comer()    { h.comiendo = true }
// func (h *hombre) pensar()   { h.pensando = true }
// func (h *hombre) sexo() string {
// 	if h.esHombre == true {
// 		return "Hombre"
// 	} else {
// 		return "Mujer"
// 	}

// }

// func HumanosRespirando(hu humano) {
// 	hu.respirar()
// 	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
// }

/* ------------------------------------------------*/
/*Genero animal */
/*
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) Escarnivoro() bool { return p.carnivoro }
func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y toy respirando")
}
func AnimalesCarnivoros(an animal) int {
	if an.Escarnivoro() == true {
		return 1
	}
	return 0
}
func main() {
	totalCanivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)
	totalCanivoros = +AnimalesCarnivoros(Dogo)
	fmt.Printf("Total carnivoros %d", totalCanivoros)
}*/

//CASO 4package main

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}
type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

/* Genero Humano */
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool { return h.vivo }
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/* ------------------------------------*/
/* Genero Animal */

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/* SerVivo  */

func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)

	fmt.Printf("Estoy vivo = %t", estoyVivo(Dogo))

}
