package mail

import (
	"net/smtp"

	"github.com/MishaNiki/go-login-signup/internal/app/model"
)

// UserRepository ...
type UserMail struct {
	mail *Mail
}

// SendingURL ...
// В будущем сдесь должен быть шаблон нормального письма в которое будут данные
func (um *UserMail) SendingURL(u *model.User, url string) error {
	err := smtp.SendMail(
		um.mail.Host+um.mail.Port,
		um.mail.Auth,
		u.Email,
		[]string{u.Email},
		[]byte("From: MyChat"+um.mail.Email+"\r\n"+
			"To:"+u.Email+"\r\n"+
			"Subject: Registration in MyChat\r\n"+
			"\r\n"+
			"To verify the registration on MyChat, follow the link:"+url+"\r\n"),
	)
	if err != nil {
		return err
	}

	return nil
}
