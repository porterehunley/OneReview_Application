package app

import(
	"net/http"
	"fmt"
	"os"
)

//TODO: Make this actaully authenticate users lmao
//It creates a middlware to intercept requests and verify them
//Notice how it takes in a handler and returns a handler, this means we can chain them arbitrarily 
//This also means we can register the middleware with the mux provided by the http package
var JwtAuthentication = func(next http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr,"User looks good to me")
		//Literally add any logic

		next.ServeHTTP(w, r)
	});
}