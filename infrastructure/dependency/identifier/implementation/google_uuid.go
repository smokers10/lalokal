package implementation

import (
	"lalokal/infrastructure/dependency/identifier"

	"github.com/google/uuid"
)

type googleUUID struct{}

// MakeIdentifier implements identifier.Contract
func (*googleUUID) MakeIdentifier() (id string) {
	rand_id, _ := uuid.NewRandom()
	return rand_id.String()
}

func GoogleUUID() identifier.Contract {
	return &googleUUID{}
}
