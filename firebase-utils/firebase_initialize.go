package firebaseutils

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	app       *firebase.App
	fcmClient *messaging.Client
}

func FirebaseInstialize() (*FirebaseApp, error) {
	opt := option.WithCredentialsFile("firebase-admin-sdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {

		return nil, fmt.Errorf("error initializing app: %v", err)

	}
	return &FirebaseApp{
		app: app,
	}, nil

}
