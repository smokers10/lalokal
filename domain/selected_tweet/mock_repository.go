package selected_tweet

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Insert(data *SelectedTweet) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) Delete(selected_tweet_id string) (failure error) {
	args := m.Mock.Called(selected_tweet_id)
	return args.Error(0)
}

func (m *MockRepository) FindByBlastingSessionId(blasting_session_id string) (result []SelectedTweet) {
	args := m.Mock.Called(blasting_session_id)
	return args.Get(0).([]SelectedTweet)
}
