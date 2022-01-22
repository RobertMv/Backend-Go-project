package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/restaurants.html", restaurantHandler)
	router.HandleFunc("/employees.html", employeeHandler)
	router.HandleFunc("/positions.html", positionHandler)
	router.HandleFunc("/", mainMenuHandler)
	http.Handle("/", router)

	err := http.ListenAndServe(":2222", nil)
	if err != nil {
		return
	}
}

func mainMenuHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/menu.html")
}

func positionHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Position page")
	if err != nil {
		return
	}
}

func restaurantHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Restaurant Page")
	if err != nil {
		return
	}
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Employee Page")
	if err != nil {
		return
	}
}
