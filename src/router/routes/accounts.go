package routes

import (
	"net/http"
	"payment_app/src/controllers"
)

var accountRoutes = []route{
	{
		uri:     "/accounts",
		method:  http.MethodPost,
		handler: controllers.CreateAccount,
	},
	{
		uri:     "/accounts",
		method:  http.MethodGet,
		handler: controllers.GetAccount,
	},
}
