package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func init() {
	// Path to your service account key JSON file
	serviceAccountKey := "/etc/secrets/serviceAccountKey.json"
	opt := option.WithCredentialsFile(serviceAccountKey)

	// Initialize the Firebase app
	var err error
	App, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}
	fmt.Println("Connected to the firebase Server!")
}
