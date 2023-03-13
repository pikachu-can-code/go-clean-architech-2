package mailprovider

type MailProvider interface {
	Send(body, subject string, receivers []string)
	SendWithAttachment(body, subject, fileName string, receivers []string)
}
