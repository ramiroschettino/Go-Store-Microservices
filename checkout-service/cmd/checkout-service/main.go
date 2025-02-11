package main

import (
	"checkout-service/internal/infrastructure/db"
	"log"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos %v", err)
	}
	log.Printf("Conexión a la base de datos exitosa")

}
