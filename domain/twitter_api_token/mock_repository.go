package twitter_api_token

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (m *MockRepository) Upsert(data *TwitterAPIToken) (failure error) {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockRepository) FindOneByTopicId(topic_id string) (result *TwitterAPIToken) {
	args := m.Mock.Called(topic_id)
	return args.Get(0).(*TwitterAPIToken)
}
