package blasting_log

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Insert(data *BlastingLogDomain) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) FindByTopicId(topic_id string) (result []BlastingLogDomain) {
	args := m.Mock.Called(topic_id)
	return args.Get(0).([]BlastingLogDomain)
}
