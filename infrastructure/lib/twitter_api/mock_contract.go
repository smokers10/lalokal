package twitter_http_request

import (
	"lalokal/domain/twitter_api_token"

	"github.com/stretchr/testify/mock"
)

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Search(keyword string, token string) (scraped_tweet *RetrunValue, failure error) {
	args := m.Mock.Called(keyword, token)
	return args.Get(0).(*RetrunValue), args.Error(1)
}

func (m *MockContract) DirectMessage(token twitter_api_token.TwitterAPIToken, event_object EOMap) (DSR *DMSuccessResponse, DER *DMErrorResponse) {
	args := m.Mock.Called(token, event_object)
	return args.Get(0).(*DMSuccessResponse), args.Get(1).(*DMErrorResponse)
}
