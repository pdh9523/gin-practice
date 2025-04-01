package email

import "gopkg.in/gomail.v2"

type GoEmailSender struct {
	from   string
	dialer *gomail.Dialer
}

func NewGoEmailSender(from, host, username, password string, port int) EmailSender {
	d := gomail.NewDialer(host, port, username, password)
	return &GoEmailSender{from: from, dialer: d}
}

func (e *GoEmailSender) Send(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return e.dialer.DialAndSend(m)
}
