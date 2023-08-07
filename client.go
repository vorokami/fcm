package fcm

import (
	"context"
	"errors"

	"firebase.google.com/go/messaging"
	"github.com/vorokami/fcm/pkg/firebase"
)

var (
	ErrInitFirebase = errors.New("firebase.InitFirebase(firebaseAuthKey) error")
	ErrGetMessaging = errors.New("firebase.GetMessaging() error")
)


type Client struct {
	fcmClient *messaging.Client
}

// NewClient creates new Firebase Cloud Messaging Client based on API key
func NewClient(firebaseAuthKey map[string]string) (*Client, error) {
	// FCM
	_, err := firebase.InitFirebase(firebaseAuthKey)
	// if Firebase is not working - exit
	if err != nil {
		return nil, ErrInitFirebase
	}

	FcmClient, err := firebase.GetMessaging()
	if err != nil {
		return nil, ErrGetMessaging
	}

	return &Client{
	fcmClient: FcmClient,
}, err
}

type SendPush interface {
		SendPushNotification(context.Context, *PushNotification) error

		// initiation of mass distribution of notifications without badges
		SendMassPushNotification(*PushNotification) error
	}

func (c *Client) SendPushNotification(ctx context.Context, pushNotification *PushNotification) error {

	tokensToSend := chunkSlice(pushNotification.Tokens, 500)

	for _, tokens := range tokensToSend {
		message := messaging.MulticastMessage{
			Data: map[string]string{
				"url": "/tabs/notifications",
			},
			Notification: &messaging.Notification{
				Title:    pushNotification.Title,
				Body:     pushNotification.Body,
				ImageURL: pushNotification.AndroidPushNotificationImage,
			},
			Android: &messaging.AndroidConfig{
				Notification: &messaging.AndroidNotification{
					Icon:              "fcm_push_icon",
					Color:             "#f45342",
					ClickAction:       "FCM_PLUGIN_ACTIVITY",
					NotificationCount: &pushNotification.Badge,
					Sound:             "sword",
					ChannelID:         "notification",
				},
				RestrictedPackageName: "ru.sami.app",
				Data: map[string]string{
					"url": "/tabs/notifications",
				},
			},
			APNS: &messaging.APNSConfig{
				Headers: map[string]string{
					"apns-priority": "10",
				},
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						Badge: &pushNotification.Badge,
						Sound: "default",
						CustomData: map[string]interface{}{
							"url": "/tabs/notifications",
						},
					},
				},
				FCMOptions: &messaging.APNSFCMOptions{
					ImageURL: "",
				},
			},
			Tokens: tokens,
		}

		_, err := c.fcmClient.SendMulticast(ctx, &message)
		if err != nil {
			return err
		}
	}
		return nil
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}