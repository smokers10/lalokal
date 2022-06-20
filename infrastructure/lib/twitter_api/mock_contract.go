package twitter_http_request

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Search(keyword string, token string) (scraped_tweet *RetrunValue, failure error) {
	args := m.Mock.Called(keyword, token)
	return args.Get(0).(*RetrunValue), args.Error(1)
}

func (m *MockContract) DirectMessage(author_id string, message string, token string) (failure error) {
	args := m.Mock.Called(author_id, message, token)
	return args.Error(0)
}
