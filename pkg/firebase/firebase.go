package firebase

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App

func InitFirebase(firebaseAuthKey map[string]string) (*firebase.App, error) {

	opts, err := json.Marshal(firebaseAuthKey)
	if err != nil {
		return nil, err
	}

	sa := option.WithCredentialsJSON(opts)
	app, err = firebase.NewApp(context.TODO(), nil, sa)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func GetApp() *firebase.App {
	return app
}
