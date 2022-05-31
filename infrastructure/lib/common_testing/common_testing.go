package common_testing

import (
	"lalokal/domain/http_response"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Expectation struct {
	Message string
	Success bool
	Status  int
	Token   string
	Data    interface{}
}

type Options struct {
	DataNotEmpty  bool
	TokenNotEmpty bool
}

var DefaultOption = &Options{
	DataNotEmpty:  false,
	TokenNotEmpty: false,
}

func Assertion(t *testing.T, expected Expectation, actual *http_response.Response, options *Options) {
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.Message, actual.Message)
	assert.Equal(t, expected.Status, actual.Status)
	assert.Equal(t, expected.Status, actual.Status)

	// assert not empty on data
	if options.DataNotEmpty {
		assert.NotEmpty(t, actual.Data)
	}

	// assert when token is required
	if options.TokenNotEmpty {
		assert.NotEmpty(t, actual.Token)
	}
}
