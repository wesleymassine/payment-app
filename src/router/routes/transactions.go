package routes

import (
	"net/http"
	"payment_app/src/controllers"
)

var transactionRoute = route{
	uri:     "/transactions",
	method:  http.MethodPost,
	handler: controllers.CreateTransaction,
}

var transactionsRoute = route{
	uri:     "/transactions",
	method:  http.MethodGet,
	handler: controllers.GetTransaction,
}
