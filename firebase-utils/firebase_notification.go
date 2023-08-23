package firebaseutils

import (
	"context"

	"firebase.google.com/go/messaging"
)

type NotificationResponse struct {
	Title               string              `json:"title"`
	Body                string              `json:"body"`
	NotificationPayload NotificationPayload `json:"payload,omitempty"`
}
type NotificationPayload struct {
}

func (firebaseapp *FirebaseApp) InitializeFcmClient() (messaging.Client, error) {
	fcmClient, err := firebaseapp.app.Messaging(context.Background())

	if err != nil {
		return *&messaging.Client{}, err
	}

	return *fcmClient, err
}

func (firebaseapp *FirebaseApp) SendSingleMessage(token string, response NotificationResponse) error {

	return nil
}
