package hotconfig

//SlackConfig A Slack or similar configuraration.
type SlackConfig struct {
	ServerAddress string
	ServerPort    int
	WebhookURL    string
	Username      string
	Password      string
}
