package mail

import "errors"

var ErrConfig = errors.New("empty config")

type Config struct {
	user     string
	password string
	smtpHost string
	smtpPort string
}

func NewConfig(user, password, host, port string) (*Config, error) {
	if isEmpty(user, password, host, port) {
		return nil, ErrConfig
	}

	return &Config{
		user:     user,
		password: password,
		smtpPort: port,
		smtpHost: host,
	}, nil
}

func isEmpty(s ...string) bool {
	for _, v := range s {
		if v == "" {
			return true
		}
	}

	return false
}
