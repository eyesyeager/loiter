package utils

import (
	"errors"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendEmail 发送邮件
func SendEmail(subject string, to []string, cc []string, text string, html string,
	from string, addr string, identity string, username string, password string, host string) error {
	e := email.NewEmail()
	e.Subject = subject
	e.From = from + " <" + username + ">"
	if len(to) == 0 {
		return errors.New("email address cannot be empty")
	}
	e.To = to
	if len(cc) != 0 {
		e.Cc = cc
	}
	if text == "" && html == "" {
		return errors.New("the email content cannot be empty")
	}
	if text != "" {
		e.Text = []byte(text)
	}
	if html != "" {
		e.HTML = []byte(html)
	}
	return e.Send(addr, smtp.PlainAuth(identity, username, password, host))
}
