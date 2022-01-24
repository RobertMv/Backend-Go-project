package router

import (
	"awesomeProject/services"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/positions/get/{id}", services.GetPosition).Methods("GET", "OPTIONS")
	router.HandleFunc("/positions/get/all", services.GetAllPositions).Methods("GET", "OPTIONS")
	router.HandleFunc("/positions/delete/{id}", services.DeletePosition).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/positions/update/{id}", services.UpdatePosition).Methods("PUT", "OPTIONS")
	router.HandleFunc("/positions/create", services.CreatePosition).Methods("POST", "OPTIONS")

	return router
}
