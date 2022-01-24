package main

import (
	"awesomeProject/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 18...")
	log.Fatal(http.ListenAndServe(":18", r))
}
