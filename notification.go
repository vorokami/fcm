package fcm

type PushNotification struct {
	Title                        string
	Body                         string
	Category                     string
	InAppImage                   string
	Tokens                       []string
	URL                          string
	ExternalURL                  string
	AndroidPushNotificationImage string
	Badge                        int
}