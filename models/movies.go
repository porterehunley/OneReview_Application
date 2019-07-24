package models

import (
	//"encoding/json"
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


