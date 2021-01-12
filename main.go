package main

import (
	"fmt"
	"log"
	"net/http"
	"payment_app/src/router"
	"payment_app/src/utils"
)

func main() {
	fmt.Println("Application Started listening on port: 5000")
	r := router.Generante()
	utils.InitializesState()
	log.Fatal(http.ListenAndServe(":5000", r))
}
