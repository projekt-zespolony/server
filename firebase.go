package main

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func firebaseNotify(notification *Notification) error {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Body,
		},
		Topic: notification.Topic,
	}

	_, err = client.Send(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}
