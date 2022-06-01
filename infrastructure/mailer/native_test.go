package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	err := NativeSMTP().Send([]string{"solarislight80@gmail.com"}, "good morning!", "hey good morning budy!")
	assert.Empty(t, err)
}
