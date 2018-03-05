package main

import (
	"flag"
	"log"

	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migracion...")
		migration.Migrate()
		log.Println("Finalizo la migracion")
	}

}
