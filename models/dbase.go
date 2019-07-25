package models

//Fun fact you can import things as something else like in python 
import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"
	"google.golang.org/api/option"
	firestore "cloud.google.com/go/firestore"
)

var dbApp *firebase.App
var db *firestore.Client


//Using defualt application credentials 
//May need to ramp this up for server deployment
func init() {
	//Load the .env 
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	dbApp = app
	db = client

	//defer client.Close()
}

//A simple handler function for the database to get a return 
func GetDBApp() *firebase.App {
	return dbApp
}

func GetDB() *firestore.Client {
	return db
}



