package mail

type Config struct {
	user     string
	password string
	smtpHost string
	smtpPort string
}

func NewConfig(user, password, host, port string) *Config {
	return &Config{
		user:     user,
		password: password,
		smtpPort: port,
		smtpHost: host,
	}
}
