package main

import (
	"/app"
	"OneReview_Application/controllers"
	"fmt"
	"net/http"
	"regexp"
)

var validAPIPath = regexp.MustCompile("^/(api)/([a-zA-Z0-9]+)$")


//Making a method to make the api handlers based on the url used
func MakeAPIHandler(fn func (http.RepsonseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := validAPIPath.FindStringSubmatch(r.URL.Path)
		fmt.Sprintf(path)
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
	//NO!
	// router := mux.NewRouter()

	http.HandleFunc("/api/", app.JwtAuthentication(MakeAPIHandler(app.GetMovie)))
	
}