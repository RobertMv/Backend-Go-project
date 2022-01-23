package main

import (
	"awesomeProject/services"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the HTTP server on port 18")
	services.CreateConnection()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/position/create", createPosition).Methods("POST")
	log.Fatal(http.ListenAndServe(":18", router))
}

func createPosition(w http.ResponseWriter, r *http.Request) {
	//requestBody, _ := ioutil.ReadAll(r.Body)
	////fmt.Println(requestBody)
	//var position models.Position
	//json.Unmarshal(requestBody, &position)
	//
	//database.Connector.Create(position)
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(position)
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
