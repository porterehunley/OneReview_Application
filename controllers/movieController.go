package controllers

import (
	"OneReview_Application/models"
	"OneReview_Application/utils"
	"net/http"
	//"fmt"
	//"encoding/json"
)


//A MUX for the api request, redirects to the correct handler based on API call
func MovieController(w http.ResponseWriter, r *http.Request, title string) {
	switch r.Method {
	case http.MethodGet:
		getMovie(w, r, title)
		return
	case http.MethodPost:
		postMovie(w, r)
		return
	case http.MethodPut:
		//TODO: Make put controller
		return
	case http.MethodDelete:
		//TODO: Make delete controller
		return 
	default:
		//TODO: give error message
		return 
	}

}

func getMovie(w http.ResponseWriter, r *http.Request, title string) {
	newMovie := &models.Movie{} 
	newMovie.Title = "testForAPIMUXOne"
	newMovie.Score = "69"
	newMovie.Date = "02/04/1999"

	//Respond with JSON
	response := utils.Message(true, "Movie found")
	response["movie"] = newMovie
	utils.Respond(w, response)

}


//TODO: Take actual data from a post request, right now this just creates a movie
func postMovie(w http.ResponseWriter, r *http.Request) {
	//Creating a movie to test firestore
	//Interesting syntax here
	newMovie := &models.Movie{} 
	newMovie.Title = "testForAPIMUXOne"
	newMovie.Score = "69"
	newMovie.Date = "02/04/1999"

	//Add movie to firestore
	response := newMovie.Create()

	utils.Respond(w, response)
}