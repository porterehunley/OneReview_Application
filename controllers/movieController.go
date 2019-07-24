package controllers

import (
	"OneReview_Application/models"
	"net/http"
	"fmt"
	"encoding/json"
)


//TODO: make this bad boy return an error when bad things happen
func GetMovie(w http.ResponseWriter, r *http.Request, title string) {
	var newMovie models.Movie 
	newMovie.Title = title
	newMovie.Score = 69
	newMovie.Date = "02/04/1999"

	//For pretty print
	//Parsing out JSON data is called "Marshaling"
	data, _ := json.MarshalIndent(newMovie, "", "  ")

	fmt.Fprintf(w,  "%s", data)
}