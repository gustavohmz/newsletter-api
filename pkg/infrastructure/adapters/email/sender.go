package email

import (
	"crypto/tls"
	"encoding/base64"
	"io"
	"os"
	"strconv"

	domain "newsletter-app/pkg/domain/models"

	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	Send(subject, body string, to []string, attachments []*domain.Attachment) error
}

type MailerSendEmailSender struct{}

func NewMailerSendEmailSender() EmailSender {
	return &MailerSendEmailSender{}
}

func (m *MailerSendEmailSender) Send(subject, body string, to []string, attachments []*domain.Attachment) error {
	emailSender := os.Getenv("emailSender")
	emailPass := os.Getenv("emailPass")
	smtpServer := os.Getenv("smtpServer")
	smtpPort := os.Getenv("smtpPort")

	smtpPortInt, err := strconv.Atoi(smtpPort)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(smtpServer, smtpPortInt, emailSender, emailPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailSender)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	for _, attachment := range attachments {
		data, err := base64.StdEncoding.DecodeString(attachment.Data)
		if err != nil {
			return err
		}

		mailer.Attach(attachment.Name, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(data)
			return err
		}))
	}

	return d.DialAndSend(mailer)
}
