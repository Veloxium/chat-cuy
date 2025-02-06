package database

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
)

var AuthClient *auth.Client

func InitFirebase() {
	opt := option.WithCredentialsFile("secret/admin-sdk.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("failed to inisialize firebase app : %v", err)
	}

	AuthClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("failed to inisialize firebase auth client : %v", err)
	}
}
