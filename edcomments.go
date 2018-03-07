package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/migration"
	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/routes"
	"github.com/urfave/negroni"
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

	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares

	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	log.Println("Iniciando el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")
}
