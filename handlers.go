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

	// get playlist info
	//playlist := getPlaylistInfo(id)
	// get playlist items
	// get durations for songs
	// make youtube request
	//list := getPlaylistItems(id)
	//fmt.Printf("list = %+v\n", list)

	// convert playlist to podcast feed, and return feed xml
	fmt.Fprintf(w, "Return playlist: %s", id)
}
