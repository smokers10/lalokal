package google_uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContract(t *testing.T) {
	GoogleUUID()
}

func TestMakeID(t *testing.T) {
	id := GoogleUUID().MakeIdentifier()
	t.Log(id)
	assert.NotEmpty(t, id)
}
