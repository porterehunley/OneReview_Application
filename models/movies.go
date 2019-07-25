package models

import (
	//"encoding/json"
	"golang.org/x/net/context"
	util "OneReview_Application/utils"
	"log"

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
	_, _, err := GetDB().Collection("movies").Add(context.Background(), map[string]interface{}{
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




