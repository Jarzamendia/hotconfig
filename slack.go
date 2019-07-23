package hotconfig

//SlackConfig A Slack or similar configuraration.
type SlackConfig struct {
	ServerAddress string
	ServerPort    string
	WebhookURL    string
	Username      string
	Password      string
	IsNull        bool
}
