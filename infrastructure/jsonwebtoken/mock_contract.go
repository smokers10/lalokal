package jsonwebtoken

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Sign(payload map[string]interface{}) (token string, failure error) {
	args := m.Mock.Called(payload)
	return args.String(0), args.Error(1)
}

func (m *MockContract) ParseToken(token string) (payload map[string]interface{}, failure error) {
	args := m.Mock.Called(token)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}
