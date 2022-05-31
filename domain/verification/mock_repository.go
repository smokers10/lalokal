package verification

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Upsert(data *Verification) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) UpdateStatus(verification_id string) (failure error) {
	args := m.Mock.Called(verification_id)
	return args.Error(0)
}

func (m *MockRepository) FindOneByEmail(email string) (result *Verification) {
	args := m.Mock.Called(email)
	return args.Get(0).(*Verification)
}
