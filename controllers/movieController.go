package controllers

import (
	"net/http"
	"fmt"
)


//TODO: make this bad boy return an error when bad things happen
func GetMovie(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Fprintf(w, "the title of this movie is %s", title)
}