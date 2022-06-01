package mailer

import (
	"fmt"
	"lalokal/infrastructure/configuration"
	"net/smtp"
)

type nativeSMTP struct {
	smtpConfiguration configuration.SMTP
}

func NativeSMTP() Contract {
	configuraton := configuration.ReadConfiguration()
	return &nativeSMTP{smtpConfiguration: configuraton.SMTP}
}

func (n *nativeSMTP) Send(reciever []string, subject string, template string) error {
	// set required data
	address := fmt.Sprintf("%s:%d", n.smtpConfiguration.Host, n.smtpConfiguration.Port)
	authentication := smtp.PlainAuth("", n.smtpConfiguration.Username, n.smtpConfiguration.Password, n.smtpConfiguration.Host)

	// email construction
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from := fmt.Sprintf("From: %s\n", n.smtpConfiguration.Sender)
	mail_subject := fmt.Sprintf("Subject: %s \n", subject)
	message := []byte(from + mail_subject + mime + template)

	// send email process
	if err := smtp.SendMail(address, authentication, n.smtpConfiguration.Username, reciever, message); err != nil {
		return err
	}

	return nil
}
