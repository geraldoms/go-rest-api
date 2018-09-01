package routes

import (
	"github.com/gorilla/mux"
)

func SetToDoRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/todos", nil).Methods("GET")
	router.HandleFunc("/api/todos/{id}", nil).Methods("GET")
	router.HandleFunc("/api/todos", nil).Methods("POST")
	router.HandleFunc("/api/todos/{id}", nil).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", nil).Methods("DELETE")

	return router
}
