// Package main provides ...
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAsset(w http.ResponseWriter, r *http.Request) {
	// extract user id
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "Return asset: %s", id)
}
