package encryption

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Hash(plain_text string) (hashed_string string) {
	args := m.Mock.Called(plain_text)
	return args.String(0)
}

func (m *MockContract) Compare(hashed_text string, plain_text string) (is_correct bool) {
	args := m.Mock.Called(hashed_text, plain_text)
	return args.Bool(0)
}
