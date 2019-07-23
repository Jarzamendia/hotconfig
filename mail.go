package hotconfig

//MailConfig A SMTP mail service
type MailConfig struct {
	ServerAddress string
	ServerPort    int
	Password      string
	Username      string
	IsSSL         bool
	CA            string
	Cert          string
	Key           string
}
