package mail

import (
	"crypto/tls"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type MailProvider struct {
	From   string
	Pass   string
	Server string
	Port   int
}

func NewMailProvider(pass string) *MailProvider {
	return &MailProvider{
		From:   "test@gmail.com",
		Server: "smtp.gmail.com",
		Port:   587,
		Pass:   pass,
	}
}

func (mail *MailProvider) Send(body, subject string, receivers []string) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", mail.From, "Monorevo")
	m.SetHeader("To", receivers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mail.Server, mail.Port, mail.From, mail.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Printf("send mail error: %v \n", err)
	}
}

func (mail *MailProvider) SendWithAttachment(body, subject, fileName string, receivers []string) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", mail.From, "Monorevo")
	m.SetHeader("To", receivers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	m.Attach(fileName)

	d := gomail.NewDialer(mail.Server, mail.Port, mail.From, mail.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
