package mail

// Config ...
type Config struct {
	Email    string `json:"email"`
	Password string `json:"password_email"`
	Host     string `json:"host_mail"`
	Port     string `json:"post_mail"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
