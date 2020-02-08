package mail

import (
	"net/smtp"
)

// Mail ...
type Mail struct {
	Auth     smtp.Auth
	Host     string
	Port     string
	Email    string
	userMail *UserMail
}

// New ...
func New(config *Config) *Mail {
	return &Mail{
		Auth: smtp.PlainAuth(
			"",
			config.Email,
			config.Password,
			config.Host,
		),
		Host:  config.Host,
		Port:  config.Port,
		Email: config.Email,
	}
}

// Ping ...
func (mail *Mail) Ping() error {

	err := smtp.SendMail(
		mail.Host+mail.Port,
		mail.Auth,
		mail.Email,
		[]string{mail.Email},
		[]byte("From: Ping <"+mail.Email+"> \r\n"+
			"To:"+mail.Email+"\r\n"+
			"Subject: Mail Check\r\n"),
	)
	if err != nil {
		return err
	}
	return nil
}

// User ...
// templete : mail.User().SendingURL('', '')
func (mail *Mail) User() *UserMail {
	if mail.userMail != nil {
		return mail.userMail
	}

	mail.userMail = &UserMail{
		mail: mail,
	}

	return mail.userMail
}
