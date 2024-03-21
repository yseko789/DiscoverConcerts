package main

import (
	"fmt"
	"net/http"
	"strings"
)

// called when user is typing artist name on the frontend, returns a list of artist name suggestions
func (app *application) searchArtistNameHandler(w http.ResponseWriter, r *http.Request) {
	keyword := app.readKeywordParam(r)
	keyword = strings.ReplaceAll(keyword, " ", "+")

	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist")

}

// called when user chooses an artist; their id is used to find similar artists
func (app *application) searchArtistSimilarHandler(w http.ResponseWriter, r *http.Request) {

}
