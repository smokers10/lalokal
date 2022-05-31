package keyword

import (
	"errors"
	"lalokal/domain/keyword"
	"lalokal/infrastructure/lib/common_testing"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	keywordRepo = keyword.MockRepository{Mock: mock.Mock{}}
	service     = keywordService{keywordRepo: &keywordRepo}
)

func TestService(t *testing.T) {
	s := KeywordService(&service.keywordRepo)
	assert.NotEmpty(t, s)
}

func TestStore(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		table := []struct {
			label    string
			input    keyword.Keyword
			expected common_testing.Expectation
		}{
			{
				label: "empty keyword",
				input: keyword.Keyword{
					Keyword: "",
					TopicId: mock.Anything,
				},
				expected: common_testing.Expectation{
					Message: "kata kunci tidak boleh kosong",
					Status:  400,
				},
			},
			{
				label: "empty topic id",
				input: keyword.Keyword{
					Keyword: mock.Anything,
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

	t.Run("failed to store keyword", func(t *testing.T) {
		input := keyword.Keyword{
			Keyword: mock.Anything,
			TopicId: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kesalahan ketika menyimpan kata kunci",
			Status:  500,
		}

		keywordRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		res := service.Store(&input)
		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		input := keyword.Keyword{
			Keyword: mock.Anything,
			TopicId: mock.Anything,
		}
		expected := common_testing.Expectation{
			Message: "kata kunci tersimpan",
			Success: true,
			Status:  200,
		}

		keywordRepo.Mock.On("Insert", mock.Anything).Return(mock.Anything, nil).Once()

		res := service.Store(&input)
		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})
}

func TestDelete(t *testing.T) {
	t.Run("empty keyword id", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "id tidak boleh kosong",
			Status:  400,
		}

		res := service.Delete("")

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to delete", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kesalahan saat menghapus kata kunci",
			Status:  500,
		}

		keywordRepo.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		res := service.Delete(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
	})

	t.Run("failed to delete", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kata kunci terhpaus",
			Success: true,
			Status:  200,
		}

		keywordRepo.Mock.On("Delete", mock.Anything).Return(nil).Once()

		res := service.Delete(mock.Anything)

		common_testing.Assertion(t, expected, res, common_testing.DefaultOption)
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

	t.Run("keyword retrieved", func(t *testing.T) {
		expected := common_testing.Expectation{
			Message: "kata kunci berhasil diambil",
			Success: true,
			Status:  200,
		}

		keywordRepo.Mock.On("FindByTopicId", mock.Anything).Return([]keyword.Keyword{
			{
				Id:      mock.Anything,
				Keyword: mock.Anything,
				TopicId: mock.Anything,
			},
		}).Once()

		res := service.ReadAll(mock.Anything)

		common_testing.Assertion(t, expected, res, &common_testing.Options{DataNotEmpty: true})
	})
}
