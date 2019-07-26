package controllers

import (
	"OneReview_Application/models"
	"OneReview_Application/utils"
	"net/http"
	//"fmt"
	"encoding/json"
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
		//Post also overwrites the datat on firebase so it doesn't really matter
		postMovie(w, r)
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
	//handeling errors?
	if title == "all" {
		response := models.GrabAllMovies()
		utils.Respond(w, response)
		return
	}
	response := models.GrabMovies(title)
	utils.Respond(w, response)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	
}


func postMovie(w http.ResponseWriter, r *http.Request) {
	//Creating a movie to test firestore
	//Interesting syntax here
	movie := &models.Movie{}
	err := json.NewDecoder(r.Body).Decode(movie)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	//Add movie to firestore
	response := movie.Create()

	utils.Respond(w, response)
}