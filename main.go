package main

import (
	"loanengine.com/mod/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRouter()
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
