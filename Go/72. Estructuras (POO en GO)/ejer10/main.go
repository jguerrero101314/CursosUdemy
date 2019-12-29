package main

import (
	"fmt"
	"time"

	us "./user"
)

type usuario struct {
	us.Usuario
}

func main() {
	u := new(usuario)
	u.AltaUsuario(1, "Joel Guerrero", time.Now(), true)

	fmt.Println(u.Usuario)
}
