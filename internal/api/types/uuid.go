package types

import (
	"github.com/google/uuid"
)

type UUID struct {
	Value string
}

func NewUUID(value string) UUID {
	return UUID{
		Value: value,
	}
}

func NewUUIDV4() UUID {
	return UUID{Value: uuid.New().String()}
}
