// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var YT_KEY string = os.Getenv("YT_KEY")

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GET", "/pl/{id:[A-Za-z0-9_-]+}", GetPlaylist},
	Route{"GET", "/asset/{id:[A-Za-z0-9_-]+}", GetAsset},
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

func init() {
	// ensure we have a yt api key from an env var
	if YT_KEY == "" {
		panic("Need youtube api key")
	}
	fmt.Printf("YT_KEY = %+v\n", YT_KEY)
}

func main() {
	router := NewRouter()
	fmt.Println("Server is running....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
