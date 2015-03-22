// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Config ConfigObj

type ConfigObj struct {
	YT_KEY string
}

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
	file, _ := os.Open("secrets.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Config)
	if err != nil {
		panic("Can't read config")
	}
}

func scratch() {
	plist := NewPlaylist("PLOJf220fGDZenl4x_jsukPVstJZ7-3FV-")
	fmt.Printf("plist.Snippet.Title = %+v\n", plist.Snippet.Title)
	fmt.Printf("plist.Snippet.Description = %+v\n", plist.Snippet.Description)
	plist.PopulatePlaylistItems()
	plist.GetItemsDetails()
	for _, item := range plist.PlaylistItems {
		fmt.Println(item.Snippet.Title)
		fmt.Println(item.Details.Duration)
	}
}

func main() {
	scratch()
	return
	router := NewRouter()
	fmt.Println("Server is running....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
