package models

import (
	//"encoding/json"
	"golang.org/x/net/context"

)
//This contains information on the movies that we will be using 

type Token struct {
	//TODO: Make a token for auth purposes
}

//The JSON field tags are purely for style purpose, they just remove capitalization for now
type Movie struct {
	Title string `json:"title"`
	Score uint `json:"score"`
	Date string `json:"date"`
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
		"date" : movie.Date
	}
}

//Not a pointer because we only care anout the data values
func addToDatabase(movie Movie, client string) error {

}


