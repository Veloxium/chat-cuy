package database

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client

func InitFirebase() error {
	opt := option.WithCredentialsFile("secret/admin-sdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	AuthClient, err = app.Auth(context.Background())
	return err
}
