package main

import (
	"OneReview_Application/app"
	"OneReview_Application/controllers"
	// "fmt"
	"net/http"
	"regexp"
	"log"
)

var validMoviePath = regexp.MustCompile("^/(api)/(movies)/[a-zA-Z0-9]+$")


//Making a method to make the api handlers based on the url used
//Some would say this limits the ammount of code used int eh handlers, as this is always the first step
func MakeAPIHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//We're using a regular expression to make sure we aren't getting nafarious paths 
		path := validMoviePath.FindStringSubmatch(r.URL.Path)

		if path == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, path[2])
	}
}


func main() {
	//Do I really want to use this mux package?
	//It routes every single request to the server though that's pretty nice
	//NO! We're not going to use it
	// router := mux.NewRouter()


	http.HandleFunc("/api/movies/", app.JwtAuthentication(MakeAPIHandler(controllers.MovieController)))

	log.Fatal(http.ListenAndServe(":8080", nil))
	
}