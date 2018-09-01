package main

import (
	"net/http"

	"./routes"
)

func main() {
	routes := routes.InitRoutes()
	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}
	server.ListenAndServe()
}
