package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() error {
	// Path to your service account key JSON file
	log.Println("Firebase connection started")
	serviceAccountKey := "serviceAccountKey.json"
	opt := option.WithCredentialsFile(serviceAccountKey)

	// Initialize the Firebase app
	var err error
	App, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
		return err
	}
	fmt.Println("Connected to the firebase Server!")
	return nil
}
