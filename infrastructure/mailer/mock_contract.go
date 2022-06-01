package mailer

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Send(reciever []string, subject string, template string) error {
	args := m.Mock.Called(reciever, subject, template)
	return args.Error(0)
}
