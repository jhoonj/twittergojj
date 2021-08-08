package main

import (
	"log"

	"github.com/jhoonj/twittergojj/bd"
	"github.com/jhoonj/twittergojj/handlers"
)

func main() {
	log.Println("empieza el curso")
	if bd.ChequeoConexion() == 0 {
		log.Fatal("sin conexion a la base de datos")
		return
	}

	handlers.Manejadores()

}
