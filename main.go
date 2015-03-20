// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GET", "/pl/{id:[a-z0-9]+}", GetPlaylist},
	Route{"GET", "/asset/{id:[a-z0-9]+}", GetAsset},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)

	}
	return router
}

func main() {
	router := NewRouter()
	fmt.Println("Server is running....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
