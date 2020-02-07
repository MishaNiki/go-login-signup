package storage

// Config ...
type Config struct {
	DataBaseURL string `json:"DB_URL"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
