package currencyapi

type Config struct {
	apiKey  string
	baseURL string
}

func NewConfig(apiKey, baseURL string) *Config {
	return &Config{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}
