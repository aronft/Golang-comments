package routes

import (
	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetUserRouter ruta para el registro de ususario
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			// negroni.HandlerFunc(controllers.ValidateToken)
			negroni.Wrap(subRouter),
		),
	)
}
