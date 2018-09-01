package main

import (
	"fmt"
	"net/http"

	"./routes"
)

func main() {
	routes := routes.InitRoutes()
	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}
	fmt.Println("Running...")
	server.ListenAndServe()
}
