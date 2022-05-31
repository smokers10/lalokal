package topic

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Insert(data *Topic) (inserted_id string, failure error) {
	args := m.Mock.Called(data)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) Update(data *Topic) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) FindByUserId(user_id string) (result []Topic) {
	args := m.Mock.Called(user_id)
	return args.Get(0).([]Topic)
}

func (m *MockRepository) FindOneById(topic_id string) (result *Topic) {
	args := m.Mock.Called(topic_id)
	return args.Get(0).(*Topic)
}
