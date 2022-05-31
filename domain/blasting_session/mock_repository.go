package blasting_session

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Insert(data *BlastingSession) (inserted_id string, failure error) {
	args := m.Mock.Called(data)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) Update(data *BlastingSession) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) FindByTopicId(topic_id string) (result []BlastingSession) {
	args := m.Mock.Called(topic_id)
	return args.Get(0).([]BlastingSession)
}

func (m *MockRepository) FindById(blasting_session_id string) (result *BlastingSession) {
	args := m.Mock.Called(blasting_session_id)
	return args.Get(0).(*BlastingSession)
}

func (m *MockRepository) UpdateStatus(blasting_session_id string, status string) (failure error) {
	args := m.Mock.Called(blasting_session_id, status)
	return args.Error(0)
}
