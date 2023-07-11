package binance

type Config struct {
	baseURL string
}

func NewConfig(baseURL string) *Config {
	return &Config{
		baseURL: baseURL,
	}
}
