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
func Init() {
	//Load the .env 
	//Might fail in deployment, but that's ok because it will get the variables from the container itself
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file: ", err)
	}
	authPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	//log.Println("Auth path: ", authPath)
	opt := option.WithCredentialsFile(authPath)
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



