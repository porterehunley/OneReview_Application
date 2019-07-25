package controllers

import (
	"OneReview_Application/models"
	"OneReview_Application/utils"
	"net/http"
	//"fmt"
	//"encoding/json"
)


//TODO: make this bad boy return an error when bad things happen
func GetMovie(w http.ResponseWriter, r *http.Request, title string) {
	

}


//TODO: Take actual data from a post request, right now this just creates a movie
func PostMovie(w http.ResponseWriter, r *http.Request, title string) {
	//Creating a movie to test firestore
	//Interesting syntax here
	newMovie := &models.Movie{} 
	newMovie.Title = title
	newMovie.Score = "69"
	newMovie.Date = "02/04/1999"

	//Add movie to firestore
	response := newMovie.Create()

	utils.Respond(w, response)
}