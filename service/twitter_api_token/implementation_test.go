package twitter_api_token

import (
	"errors"
	"lalokal/domain/twitter_api_token"
	"lalokal/infrastructure/lib/common_testing"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	twitterAPIRepository = twitter_api_token.MockRepository{Mock: mock.Mock{}}
	service              = twitterAPIService{twitterAPIRepository: &twitterAPIRepository}
)

func TestTWitterAPIService(t *testing.T) {
	s := TwitterAPIService(&service.twitterAPIRepository)
	assert.NotEmpty(t, s)
}

func TestStore(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    twitter_api_token.TwitterAPIToken
			expected common_testing.Expectation
		}{
			{
				label: "empty API token",
				input: twitter_api_token.TwitterAPIToken{
					APIToken:       "",
					ConsumerKey:    mock.Anything,
					ConsumerSecret: mock.Anything,
					AccessToken:    mock.Anything,
					AccessSecret:   mock.Anything,
					TopicId:        mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "API token tidak boleh kosong",
					Status:  400,
				},
			},

			{
				label: "empty access secret",
				input: twitter_api_token.TwitterAPIToken{
					APIToken:       mock.Anything,
					ConsumerKey:    mock.Anything,
					ConsumerSecret: mock.Anything,
					AccessToken:    mock.Anything,
					AccessSecret:   "",
					TopicId:        mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "access secret tidak boleh kosong",
					Status:  400,
				},
			},

			{
				label: "empty access token",
				input: twitter_api_token.TwitterAPIToken{
					APIToken:       mock.Anything,
					ConsumerKey:    mock.Anything,
					ConsumerSecret: mock.Anything,
					AccessToken:    "",
					AccessSecret:   mock.Anything,
					TopicId:        mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "access token tidak boleh kosong",
					Status:  400,
				},
			},

			{
				label: "empty consumer key",
				input: twitter_api_token.TwitterAPIToken{
					APIToken:       mock.Anything,
					ConsumerKey:    "",
					ConsumerSecret: mock.Anything,
					AccessToken:    mock.Anything,
					AccessSecret:   mock.Anything,
					TopicId:        mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "consumer key tidak boleh kosong",
					Status:  400,
				},
			},

			{
				label: "empty API token",
				input: twitter_api_token.TwitterAPIToken{
					APIToken:       mock.Anything,
					ConsumerKey:    mock.Anything,
					ConsumerSecret: "",
					AccessToken:    mock.Anything,
					AccessSecret:   mock.Anything,
					TopicId:        mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "consumer secret tidak boleh kosong",
					Status:  400,
				},
			},

			{
				label: "empty API token",
				input: twitter_api_token.TwitterAPIToken{
					APIToken:       mock.Anything,
					ConsumerKey:    mock.Anything,
					ConsumerSecret: mock.Anything,
					AccessToken:    mock.Anything,
					AccessSecret:   mock.Anything,
					TopicId:        "",
				},
				expected: common_testing.Expectation{
					Message: "id topik tidak boleh kosong",
					Status:  400,
				},
			},
		}

		for _, tb := range table {
			res := service.Store(&tb.input)

			common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
		}
	})

	t.Run("failed to upser twitter api token", func(t *testing.T) {
		input := twitter_api_token.TwitterAPIToken{
			APIToken:       mock.Anything,
			ConsumerKey:    mock.Anything,
			ConsumerSecret: mock.Anything,
			AccessToken:    mock.Anything,
			AccessSecret:   mock.Anything,
			TopicId:        mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "kesalahan ketika menyimpan token twitter",
			Status:  500,
		}

		twitterAPIRepository.Mock.On("Upsert", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.Store(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success condition", func(t *testing.T) {
		input := twitter_api_token.TwitterAPIToken{
			APIToken:       mock.Anything,
			ConsumerKey:    mock.Anything,
			ConsumerSecret: mock.Anything,
			AccessToken:    mock.Anything,
			AccessSecret:   mock.Anything,
			TopicId:        mock.Anything,
		}

		expected := common_testing.Expectation{
			Message: "token twitter tersimpan",
			Success: true,
			Status:  200,
		}

		twitterAPIRepository.Mock.On("Upsert", mock.Anything).Return(nil).Once()

		res := service.Store(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})
}

func TestRead(t *testing.T) {
	t.Run("empty topic id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}

		res := service.Read("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("empty topic id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "token twitter berhasil diambil",
			Success: true,
			Status:  200,
		}

		result := twitter_api_token.TwitterAPIToken{
			APIToken:       mock.Anything,
			ConsumerKey:    mock.Anything,
			ConsumerSecret: mock.Anything,
			AccessToken:    mock.Anything,
			AccessSecret:   mock.Anything,
			TopicId:        mock.Anything,
		}
		twitterAPIRepository.Mock.On("FindOneByTopicId", mock.Anything).Return(&result).Once()

		res := service.Read(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}
