package router

import (
	"payment_app/src/router/routes"

	"github.com/gorilla/mux"
)

//Generante return the router with all routes configured
func Generante() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
