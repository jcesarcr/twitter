package main

import (
	"log"

	"github.com/jcesarcr/twitter/bd"
	"github.com/jcesarcr/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
