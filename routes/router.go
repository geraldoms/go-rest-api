package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Set the routes for todo model
	router = SetToDoRoutes(router)

	return router
}
