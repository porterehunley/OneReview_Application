package models

import (
	//"encoding/json"
	"golang.org/x/net/context"
	util "OneReview_Application/utils"
	"log"
	"google.golang.org/api/iterator"

)
//This contains information on the movies that we will be using 

type Token struct {
	//TODO: Make a token for auth purposes
}

//The JSON field tags are purely for style purpose, they just remove capitalization for now
type Movie struct {
	Title string `json:"title"`
	Score string `json:"score"`
	Date string `json:"date"`
}	


//To export things in a package, PUT A CAPITAL LETTER ON THE FUNCTION
func (movie *Movie) Create() (map[string] interface{}) {
	//I should probably figure out what values I just set to Nil
	//Also what exactly is an interface? Why am I getting errors when I use toKeyValue??
	_, err := GetDB().Collection("movies").Doc(movie.Title).Set(context.Background(), map[string]interface{}{
        "title" : movie.Title, 
		"score" : movie.Score,
		"date" : movie.Date,
	})
	if err != nil {
        log.Fatalf("Failed adding movie: %v", err)
	}

	//Return a response on success
	response := util.Message(true, "Movie has been created and stored")
	response["movie"] = movie
	return response
}

func GrabAllMovies() (map[string] interface{}) {
	iter := GetDB().Collection("movies").Documents(context.Background())
	var movies []map[string]interface{}

	for true {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            log.Println("Error getting all movies: ", err)
            return nil
        }
        movies = append(movies, doc.Data())
	}

	response := util.Message(true, "Movies have been found")
	response["movies"] = movies
	return response
}

func DeleteMovie(title string) (map[string] interface{}) {
	_, err := GetDB().Collection("movies").Doc(title).Delete(context.Background())
	if err != nil {
		log.Printf("Error deleting movie: %s", err)
		response := util.Message(false, "Failed to delete movie")
	    return response
	}
	response := util.Message(true, "Movie has been deleted")
	return response
}

//Get (grab) the movies from firestore based from the title of the movie
func GrabMovies(title string) (map[string] interface{}) {
	dataSnap, err := GetDB().Collection("movies").Doc(title).Get(context.Background())
	if err != nil {
	    log.Println("Failed grabbing movie: ", err)
	    response := util.Message(false, "Movie not found")
	    return response
	}

	response := util.Message(true, "Movie has been found")
	response["movies"] = dataSnap.Data()
	return response
}

func (movie *Movie) getScore() {
	//TODO: This is where we will get the score of movie from database
}

func (movie *Movie) getDate() {
	//TODO: This is where we will get the date of a movie 
}



func (movie *Movie) toKeyValue() (map[string]interface{}) {
	return map[string]interface{} {
		"title" : movie.Title, 
		"score" : movie.Score,
		"date" : movie.Date,
	}
}




