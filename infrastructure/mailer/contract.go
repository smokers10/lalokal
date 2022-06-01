package mailer

type Contract interface {
	Send(reciever []string, subject string, template string) error
}
