package user

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Insert(data *RegisterData) (inserted_id string, failure error) {
	args := m.Mock.Called(data)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) UpdatePassword(data *ResetPasswordData) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) Update(data *User) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) FindOneByEmail(email string) (result *User) {
	args := m.Mock.Called(email)
	return args.Get(0).(*User)
}

func (m *MockRepository) FindOneById(user_id string) (result *User) {
	args := m.Mock.Called(user_id)
	return args.Get(0).(*User)
}
