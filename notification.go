package fcm

type PushNotification struct {
	Title                        string // title of push notification
	Body                         string // text body of push notification
	Tokens                       []string
	URL                          string
	ExternalURL                  string
	AndroidPushNotificationImage string // image URL
	Badge                        int    // unread notification count to set on app icon
	Sound                        string // name of sound of the push notification. This sound is set when building the application
	ChannelID                    string
}
