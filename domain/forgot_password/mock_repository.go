package forgot_password

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Insert(data *ForgotPassword) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) FindOneByToken(token string) (result *ForgotPassword) {
	args := m.Mock.Called(token)
	return args.Get(0).(*ForgotPassword)
}

func (m *MockRepository) Delete(token string) (failure error) {
	args := m.Mock.Called(token)
	return args.Error(0)
}
