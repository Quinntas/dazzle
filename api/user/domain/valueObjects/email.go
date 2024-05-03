package userValueObjects

import (
	"github.com/quinntas/go-rest-template/internal/api/utils/guard"
)

type Email struct {
	Value string
}

func NewEmail(value string) Email {
	return Email{
		Value: value,
	}
}

func NewEmailWithValidation(value string) (Email, error) {
	err := guard.AgainstBadEmail("email", value)
	if err != nil {
		return Email{}, err
	}

	return Email{
		Value: value,
	}, nil
}
