// Package main provides ...
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	// extract user id
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "Return playlist: %s", id)
}
