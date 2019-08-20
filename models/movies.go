package models

import (
	// "encoding/json"
	"golang.org/x/net/context"
	util "OneReview_Application/utils"
	"log"
	"google.golang.org/api/iterator"
	"net/http"
	"bytes"
	"strings"
	"io/ioutil"
	"time"
)
//This contains information on the movies that we will be using 

type Token struct {
	//TODO: Make a token for auth purposes
}

//The JSON field tags are purely for style purpose, they just remove capitalization for now
//Omitempty means that we dont need the field to be there to create a movie
type Movie struct {
	Title string `json:"title"`
	Score string `json:"score"`
	Date string `json:"date"`
	Image string `json:"image,omitempty"`
}	


//To export things in a package, PUT A CAPITAL LETTER ON THE FUNCTION
//Adding image collection here as well
func (movie *Movie) Create() (map[string] interface{}) {
	//I should probably figure out what values I just set to Nil
	//Also what exactly is an interface? Why am I getting errors when I use toKeyValue??
	if movie.Image == "" {
		log.Println("getting image")
		movie.getImage()
	}

	_, err := GetDB().Collection("movies").Doc(movie.Title).Set(context.Background(), map[string]interface{}{
        "title" : movie.Title, 
		"score" : movie.Score,
		"date" : movie.Date,
		"image" : movie.Image,
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

func (movie *Movie) getImage() {
	client := &http.Client{
        Timeout: 30 * time.Second,
    }

    titleFields := strings.Fields(movie.Title)
    searchBuff := bytes.Buffer{}
    searchBuff.WriteString("https://www.imdb.com/find?ref_=nv_sr_fn&q=")
    for i := 0; i < len(titleFields); i++ {
    	if i > 0 {
    		searchBuff.WriteString("+")
    	}
    	searchBuff.WriteString(titleFields[i])
    }
    searchBuff.WriteString("&s=all")
    searchUrl := searchBuff.String()
    log.Println(searchUrl)

    response, err := client.Get(searchUrl)
    if err != nil {
        log.Fatal(err)
    }
    //Closes the reesposne body when the surronding function returns
    defer response.Body.Close()

    // Get the response body as a string
    dataInBytes, err := ioutil.ReadAll(response.Body)
    pageContent := string(dataInBytes)

    pageContentStartIndex := strings.Index(pageContent, "<div id=\"pagecontent\" class=\"pagecontent\">")
    if pageContentStartIndex == -1 {
        log.Println("No start page content element found")
        return
    }

    pageContent = pageContent[pageContentStartIndex:]

    urlStartIndex := strings.Index(pageContent, "https://m.media-amazon.com/images/M/")
    if urlStartIndex == -1 {
        log.Println("No start image element found")
        return
    }

   	urlEndIndex := strings.Index(pageContent, "UX32_CR0,0,32,44_AL_.jpg")
    if urlEndIndex == -1 {
        log.Println("No end image element found")
        return
    }

    imageBuff := bytes.Buffer{}
    imageBuff.WriteString(pageContent[urlStartIndex:urlEndIndex])
    imageBuff.WriteString("UX182_CR0,0,182,268_AL_.jpg")
    imageUrl := imageBuff.String()
    log.Println(imageUrl)
    movie.Image = imageUrl
}

func (movie *Movie) toKeyValue() (map[string]interface{}) {
	return map[string]interface{} {
		"title" : movie.Title, 
		"score" : movie.Score,
		"date" : movie.Date,
	}
}




