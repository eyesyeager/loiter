package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendEmailWithText 发送邮件(Text格式)
func SendEmailWithText(subject string, from string, to string, cc string, text string,
	addr string, identity string, username string, password string, host string) error {
	e := email.NewEmail()
	e.Subject = subject
	e.From = from
	e.To = []string{to}
	e.Cc = []string{cc}
	e.Text = []byte(text)
	return e.Send(addr, smtp.PlainAuth(identity, username, password, host))
}

// SendEmailWithHTML 发送邮件(HTML格式)
func SendEmailWithHTML(subject string, from string, to string, cc string, html string,
	addr string, identity string, username string, password string, host string) error {
	e := email.NewEmail()
	e.Subject = subject
	e.From = from
	e.To = []string{to}
	e.Cc = []string{cc}
	e.HTML = []byte(html)
	return e.Send(addr, smtp.PlainAuth(identity, username, password, host))
}
