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

func (m *MockRepository) Count(topic_id string) (count int) {
	args := m.Mock.Called(topic_id)
	return args.Int(0)
}

func (m *MockRepository) LogPercentage(blasting_session_id string) (total_message int, success_count int, failed_count int, success_percentage float32, fail_percentage float32) {
	args := m.Mock.Called(blasting_session_id)
	return args.Int(0), args.Int(1), args.Int(2), float32(args.Int(3)), float32(args.Int(4))
}
