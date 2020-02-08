package templates

// Config ...
type Config struct {
	Login  string `json:"pg_login"`
	SignUp string `json:"pg_signup"`
	Forgod string `json:"pg_forgod"`
	Home   string `json:"pg_home"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
