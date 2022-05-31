package keyword

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Insert(data *Keyword) (inserted_id string, failure error) {
	args := m.Mock.Called(data)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) Delete(keyword_id string) (failure error) {
	args := m.Mock.Called(keyword_id)
	return args.Error(0)
}

func (m *MockRepository) FindByTopicId(topic_id string) (result []Keyword) {
	args := m.Mock.Called(topic_id)
	return args.Get(0).([]Keyword)
}
