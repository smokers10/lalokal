package twitter_http_request

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Search(keyword string) (scraped_tweet []map[string]interface{}, failure error) {
	args := m.Mock.Called(keyword)
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func (m *MockContract) DirectMessage(author_id string, message string) (failure error) {
	args := m.Mock.Called(author_id, message)
	return args.Error(0)
}
