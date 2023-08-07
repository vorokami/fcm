package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

var cliFcm *messaging.Client

func initMessaging(app *firebase.App) error {
	ctx := context.Background()
	cli, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	cliFcm = cli

	return nil
}

func GetMessaging() (*messaging.Client, error) {
	err := initMessaging(app)
	if err != nil {
		return nil, err
	}
	return cliFcm, nil
}
