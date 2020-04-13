package firebase

import (
	"context"

	fb "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"github.com/projekt-zespolony/server/pkg/types"
)

type Firebase struct {
	app    *fb.App
	client *messaging.Client
}

type Options struct {
	Disabled bool
}

func New(options *Options) (*Firebase, error) {
	if options.Disabled {
		return &Firebase{}, nil
	}

	app, err := fb.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return &Firebase{
		app:    app,
		client: client,
	}, nil
}

func (firebase *Firebase) Notify(notification *types.Notification) error {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Body,
		},
		Topic: notification.Topic,
	}

	_, err := firebase.client.Send(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}
