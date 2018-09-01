package routes

import (
	"github.com/gorilla/mux"

	"../controllers"
)

func SetToDoRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/todos", controllers.GetAll).Methods("GET")
	router.HandleFunc("/api/todos/{id}", controllers.GetByID).Methods("GET")
	router.HandleFunc("/api/todos", controllers.Create).Methods("POST")
	//router.HandleFunc("/api/todos/{id}", nil).Methods("PUT")
	//router.HandleFunc("/api/todos/{id}", nil).Methods("DELETE")

	return router
}
