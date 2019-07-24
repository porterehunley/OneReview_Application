package utils

//Fun fact you can import things as something else like in python 
import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

//Using defualt application credentials 
//May need to ramp this up for server deployment
func InitFirebaseApp() (*firebase.App, error) {
	//Load the .env 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
