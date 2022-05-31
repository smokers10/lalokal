package blasting_session

import (
	"errors"
	"lalokal/domain/blasting_log"
	"lalokal/domain/blasting_session"
	"lalokal/domain/keyword"
	"lalokal/domain/selected_tweet"
	"lalokal/domain/twitter_api_token"
	"lalokal/infrastructure/lib/common_testing"
	twitter_http_request "lalokal/infrastructure/lib/twitter_api"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	blastingLogRepo     = blasting_log.MockRepository{Mock: mock.Mock{}}
	blastingSessionRepo = blasting_session.MockRepository{Mock: mock.Mock{}}
	twitterApiTokenRepo = twitter_api_token.MockRepository{Mock: mock.Mock{}}
	keywordRepo         = keyword.MockRepository{Mock: mock.Mock{}}
	twitter             = twitter_http_request.MockContract{Mock: mock.Mock{}}
	selectedTweetRepo   = selected_tweet.MockRepository{Mock: mock.Mock{}}
	service             = blastingSessionService{
		blastingLogRepository:     &blastingLogRepo,
		blastingSessionRepository: &blastingSessionRepo,
		twitterAPITokenRepository: &twitterApiTokenRepo,
		keywordRepository:         &keywordRepo,
		twitter:                   &twitter,
		selectedTweetRepository:   &selectedTweetRepo,
	}
)

func TestService(t *testing.T) {
	s := BlastingSessionService(&service.blastingLogRepository, &service.blastingSessionRepository, &service.twitterAPITokenRepository,
		&service.keywordRepository, &service.twitter, &service.selectedTweetRepository)

	assert.NotEmpty(t, s)
}

func TestStore(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    blasting_session.BlastingSession
			expected common_testing.Expectation
		}{
			{
				label: "empty message",
				input: blasting_session.BlastingSession{
					Title:   mock.Anything,
					Message: "",
					TopicId: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "pesan tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty title",
				input: blasting_session.BlastingSession{
					Title:   "",
					Message: mock.Anything,
					TopicId: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "judul sesi blasting tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty topic id",
				input: blasting_session.BlastingSession{
					Title:   mock.Anything,
					Message: mock.Anything,
					TopicId: "",
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

	t.Run("failed to store blasting session", func(t *testing.T) {
		input := blasting_session.BlastingSession{
			Title:   mock.Anything,
			Message: mock.Anything,
			TopicId: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kesalahan ketika menyimpan sesi blasting",
			Status:  500,
		}

		blastingSessionRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		res := service.Store(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := blasting_session.BlastingSession{
			Title:   mock.Anything,
			Message: mock.Anything,
			TopicId: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "sesi blasting tersimpan",
			Success: true,
			Status:  200,
		}

		blastingSessionRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, nil).Once()

		res := service.Store(&input)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    blasting_session.BlastingSession
			expected common_testing.Expectation
		}{
			{
				label: "empty message",
				input: blasting_session.BlastingSession{
					Id:      mock.Anything,
					Title:   mock.Anything,
					Message: "",
					TopicId: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "pesan tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty title",
				input: blasting_session.BlastingSession{
					Id:      mock.Anything,
					Title:   "",
					Message: mock.Anything,
					TopicId: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "judul sesi blasting tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty topic id",
				input: blasting_session.BlastingSession{
					Id:      mock.Anything,
					Title:   mock.Anything,
					Message: mock.Anything,
					TopicId: "",
				},
				expected: common_testing.Expectation{
					Message: "id topik tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty id",
				input: blasting_session.BlastingSession{
					Id:      "",
					Title:   mock.Anything,
					Message: mock.Anything,
					TopicId: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "id tidak boleh kosong",
					Status:  400,
				},
			},
		}

		for _, tb := range table {
			res := service.Update(&tb.input)

			common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
		}
	})

	t.Run("failed to update", func(t *testing.T) {
		input := blasting_session.BlastingSession{
			Id:      mock.Anything,
			Title:   mock.Anything,
			Message: mock.Anything,
			TopicId: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kesalahan ketika update sesi blasting",
			Status:  500,
		}

		blastingSessionRepo.Mock.On("Update", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.Update(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := blasting_session.BlastingSession{
			Id:      mock.Anything,
			Title:   mock.Anything,
			Message: mock.Anything,
			TopicId: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "sesi blasting terupdate",
			Success: true,
			Status:  200,
		}

		blastingSessionRepo.Mock.On("Update", mock.Anything).Return(nil).Once()

		res := service.Update(&input)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestReadAll(t *testing.T) {
	t.Run("empty topic id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}

		res := service.ReadAll("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("blasting session retrieved", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "sesi blasting berhasil diambil",
			Success: true,
			Status:  200,
		}

		retrieved_dummy := []blasting_session.BlastingSession{
			{
				Id:      mock.Anything,
				Title:   mock.Anything,
				Message: mock.Anything,
				Status:  mock.Anything,
				TopicId: mock.Anything,
			},
		}

		blastingSessionRepo.Mock.On("FindByTopicId", mock.Anything).Return(retrieved_dummy).Once()

		res := service.ReadAll(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestDetail(t *testing.T) {
	t.Run("empty id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id tidak boleh kosong",
			Status:  400,
		}

		res := service.Detail("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("blasting session retrieved", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "sesi blasting berhasil diambil",
			Success: true,
			Status:  200,
		}

		retrieved_dummy := &blasting_session.BlastingSession{
			Id:      mock.Anything,
			Title:   mock.Anything,
			Message: mock.Anything,
			Status:  mock.Anything,
			TopicId: mock.Anything,
		}

		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(retrieved_dummy).Once()

		res := service.Detail(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestScrape(t *testing.T) {
	t.Run("empty blasting session id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id tidak boleh kosong",
			Status:  400,
		}

		res := service.Scrape("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("empty blasting session", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "sesi blasting tidak ditemukan",
			Status:  404,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{}

		// action & asser

		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()

		res := service.Scrape(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("empty twitter token", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "api token twitter untuk topik tidak ada",
			Status:  404,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, TopicId: mock.Anything}
		twitter_api_token := &twitter_api_token.TwitterAPIToken{}

		// action & asser
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		twitterApiTokenRepo.Mock.On("FindOneByTopicId", mock.Anything).Return(twitter_api_token).Once()

		res := service.Scrape(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("empty keywords", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kata kunci kosong",
			Status:  404,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, TopicId: mock.Anything}
		twitter_api_token := &twitter_api_token.TwitterAPIToken{Token: mock.Anything, Secret: mock.Anything}
		keywords := []keyword.Keyword{}

		// action & asser
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		twitterApiTokenRepo.Mock.On("FindOneByTopicId", mock.Anything).Return(twitter_api_token).Once()
		keywordRepo.Mock.On("FindByTopicId", mock.Anything).Return(keywords).Once()

		res := service.Scrape(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to scrape tweets", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kesalahan ketika mengambil tuitan",
			Status:  500,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, TopicId: mock.Anything}
		twitter_api_token := &twitter_api_token.TwitterAPIToken{Token: mock.Anything, Secret: mock.Anything}
		keywords := []keyword.Keyword{
			{
				Id:      mock.Anything,
				Keyword: mock.Anything,
				TopicId: mock.Anything,
			},
		}

		// action & asser
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		twitterApiTokenRepo.Mock.On("FindOneByTopicId", mock.Anything).Return(twitter_api_token).Once()
		keywordRepo.Mock.On("FindByTopicId", mock.Anything).Return(keywords).Once()
		twitter.Mock.On("Search", mock.Anything).Return([]map[string]interface{}{}, errors.New(mock.Anything)).Once()

		res := service.Scrape(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "tuitan berhasil diambil",
			Success: true,
			Status:  200,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, TopicId: mock.Anything}
		twitter_api_token := &twitter_api_token.TwitterAPIToken{Token: mock.Anything, Secret: mock.Anything}
		keywords := []keyword.Keyword{
			{
				Id:      mock.Anything,
				Keyword: mock.Anything,
				TopicId: mock.Anything,
			},
		}
		scraped_tweet := []map[string]interface{}{
			{
				"text":      mock.Anything,
				"author_id": mock.Anything,
			},
		}

		// action & asser
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		twitterApiTokenRepo.Mock.On("FindOneByTopicId", mock.Anything).Return(twitter_api_token).Once()
		keywordRepo.Mock.On("FindByTopicId", mock.Anything).Return(keywords).Once()
		twitter.Mock.On("Search", mock.Anything).Return(scraped_tweet, nil).Once()

		res := service.Scrape(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestBlast(t *testing.T) {
	t.Run("empty blasting session id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id tidak boleh kosong",
			Status:  400,
		}

		res := service.Blast("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("empty blasting session", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "sesi blasting tidak ditemukan",
			Status:  404,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{}

		// action & asser

		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()

		res := service.Blast(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("revoked blasting session", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "sesi blasting sudah selesai",
			Status:  403,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, Status: "revoked"}

		// action & assert

		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()

		res := service.Blast(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("no selected tweets", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "tidak ada tuitan yang dipilih",
			Status:  404,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, Status: "waiting"}
		selected_tweets := []selected_tweet.SelectedTweet{}

		// action & assert
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		selectedTweetRepo.Mock.On("FindByBlastingSessionId", mock.Anything).Return(selected_tweets).Once()

		res := service.Blast(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to send message", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "blasting selesai",
			Success: true,
			Status:  200,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, Status: "waiting"}
		selected_tweets := []selected_tweet.SelectedTweet{
			{
				Id:                mock.Anything,
				AuthorId:          mock.Anything,
				TweetId:           mock.Anything,
				Text:              mock.Anything,
				BlastingSessionId: mock.Anything,
			},
		}

		// action & assert
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		selectedTweetRepo.Mock.On("FindByBlastingSessionId", mock.Anything).Return(selected_tweets).Once()

		twitter.Mock.On("DirectMessage", mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()
		blastingLogRepo.Mock.On("Insert", mock.Anything).Return(nil).Once()
		blastingSessionRepo.Mock.On("UpdateStatus", mock.Anything, mock.Anything).Return(nil).Once()

		res := service.Blast(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success to send message", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "blasting selesai",
			Success: true,
			Status:  200,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, Status: "waiting"}
		selected_tweets := []selected_tweet.SelectedTweet{
			{
				Id:                mock.Anything,
				AuthorId:          mock.Anything,
				TweetId:           mock.Anything,
				Text:              mock.Anything,
				BlastingSessionId: mock.Anything,
			},
		}

		// action & assert
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		selectedTweetRepo.Mock.On("FindByBlastingSessionId", mock.Anything).Return(selected_tweets).Once()

		twitter.Mock.On("DirectMessage", mock.Anything, mock.Anything).Return(nil).Once()
		blastingLogRepo.Mock.On("Insert", mock.Anything).Return(nil).Once()
		blastingSessionRepo.Mock.On("UpdateStatus", mock.Anything, mock.Anything).Return(nil).Once()

		res := service.Blast(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to update blasting session status", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "gagal update status sesi blasting",
			Status:  500,
		}

		// dummy query result
		blasting_sessions := &blasting_session.BlastingSession{Id: mock.Anything, Status: "waiting"}
		selected_tweets := []selected_tweet.SelectedTweet{
			{
				Id:                mock.Anything,
				AuthorId:          mock.Anything,
				TweetId:           mock.Anything,
				Text:              mock.Anything,
				BlastingSessionId: mock.Anything,
			},
		}

		// action & assert
		blastingSessionRepo.Mock.On("FindById", mock.Anything).Return(blasting_sessions).Once()
		selectedTweetRepo.Mock.On("FindByBlastingSessionId", mock.Anything).Return(selected_tweets).Once()

		twitter.Mock.On("DirectMessage", mock.Anything, mock.Anything).Return(nil).Once()
		blastingLogRepo.Mock.On("Insert", mock.Anything).Return(nil).Once()
		blastingSessionRepo.Mock.On("UpdateStatus", mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.Blast(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})
}
