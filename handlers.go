// Package main provides ...
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Test(w http.ResponseWriter, r *http.Request) {

	plist := Playlist{}
	file, _ := os.Open("sample.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&plist)
	if err != nil {
		panic(err)
	}

	// convert playlist to feed
	feed := NewFeed(plist)

	// return XML of feed
	w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	xml, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", string(xml))
}

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	// extract user id
	vars := mux.Vars(r)
	id := vars["id"]

	// playlist requested
	fmt.Println("Playlist requested: ", id)

	// create playlist from ID
	plist := NewPlaylist(id)      // makes API call to get playlist title, description
	plist.PopulatePlaylistItems() // gets playlist items
	plist.GetItemsDetails()       // gets durations for each item

	// convert playlist to feed
	feed := NewFeed(*plist)

	// return XML of feed
	w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := xml.NewEncoder(w).Encode(feed); err != nil {
		panic(err)
	}
}
