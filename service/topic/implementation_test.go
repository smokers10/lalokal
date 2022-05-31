package topic

import (
	"errors"
	"lalokal/domain/topic"
	"lalokal/infrastructure/lib/common_testing"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	topicRepo = topic.MockRepository{Mock: mock.Mock{}}
	service   = topicService{topicRepository: &topicRepo}
)

func TestService(t *testing.T) {
	s := TopicService(&service.topicRepository)

	assert.NotEmpty(t, s)
}

func TestStore(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    topic.Topic
			expected common_testing.Expectation
		}{
			{
				label: "empty title",
				input: topic.Topic{
					Id:          mock.Anything,
					Title:       "",
					Description: mock.Anything,
					UserId:      mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "judul topik tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty description",
				input: topic.Topic{
					Id:          mock.Anything,
					Title:       mock.Anything,
					Description: "",
					UserId:      mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "deskripsi topik tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty user id",
				input: topic.Topic{
					Id:          mock.Anything,
					Title:       mock.Anything,
					Description: mock.Anything,
					UserId:      "",
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

	t.Run("failed to insert topic", func(t *testing.T) {
		input := topic.Topic{
			Id:          mock.Anything,
			Title:       mock.Anything,
			Description: mock.Anything,
			UserId:      mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kesalahan saat menyimpan topic",
			Status:  500,
		}

		topicRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		res := service.Store(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := topic.Topic{
			Id:          mock.Anything,
			Title:       mock.Anything,
			Description: mock.Anything,
			UserId:      mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "topic tersimpan",
			Success: true,
			Status:  200,
		}

		topicRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, nil).Once()

		res := service.Store(&input)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    topic.Topic
			expected common_testing.Expectation
		}{
			{
				label: "empty title",
				input: topic.Topic{
					Id:          mock.Anything,
					Title:       "",
					Description: mock.Anything,
					UserId:      mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "judul topik tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty description",
				input: topic.Topic{
					Id:          mock.Anything,
					Title:       mock.Anything,
					Description: "",
					UserId:      mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "deskripsi topik tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty user id",
				input: topic.Topic{
					Id:          mock.Anything,
					Title:       mock.Anything,
					Description: mock.Anything,
					UserId:      "",
				},
				expected: common_testing.Expectation{
					Message: "id pengguna tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty id",
				input: topic.Topic{
					Id:          "",
					Title:       mock.Anything,
					Description: mock.Anything,
					UserId:      mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "id topik tidak boleh kosong",
					Status:  400,
				},
			},
		}

		for _, tb := range table {
			res := service.Update(&tb.input)

			common_testing.Assertion(t, tb.expected, res, common_testing.DefaultOption)
		}
	})

	t.Run("failed to update topic", func(t *testing.T) {
		input := topic.Topic{
			Id:          mock.Anything,
			Title:       mock.Anything,
			Description: mock.Anything,
			UserId:      mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kesalahan saat update topic",
			Status:  500,
		}

		topicRepo.Mock.On("Update", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.Update(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := topic.Topic{
			Id:          mock.Anything,
			Title:       mock.Anything,
			Description: mock.Anything,
			UserId:      mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "topic terupdate",
			Success: true,
			Status:  200,
		}

		topicRepo.Mock.On("Update", mock.Anything).Return(nil).Once()

		res := service.Update(&input)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})
}

func TestReadAll(t *testing.T) {
	t.Run("empty user id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id pengguna tidak boleh kosong",
			Status:  400,
		}

		res := service.ReadAll("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("retrieved topic", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "topik berhasil diambil",
			Success: true,
			Status:  200,
		}

		topicRepo.Mock.On("FindByUserId", mock.Anything).Return([]topic.Topic{
			{
				Id:          mock.Anything,
				Title:       mock.Anything,
				Description: mock.Anything,
				UserId:      mock.Anything,
			},
		})

		res := service.ReadAll(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}

func TestDetail(t *testing.T) {
	t.Run("empty topic id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}

		res := service.Detail("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("retrieved topic", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "topik berhasil diambil",
			Success: true,
			Status:  200,
		}

		topicRepo.Mock.On("FindOneById", mock.Anything).Return(&topic.Topic{
			Id:          mock.Anything,
			Title:       mock.Anything,
			Description: mock.Anything,
			UserId:      mock.Anything,
		})

		res := service.Detail(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}
