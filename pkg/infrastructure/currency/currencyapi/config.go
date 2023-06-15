package currencyapi

type Config struct {
	apiKey string
}

func NewConfig(apiKey string) *Config {
	return &Config{
		apiKey: apiKey,
	}
}
