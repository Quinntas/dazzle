package userValueObjects

import (
	"context"

	"github.com/quinntas/go-rest-template/internal/api/encryption"
	"github.com/quinntas/go-rest-template/internal/api/utils/env"
	"github.com/quinntas/go-rest-template/internal/api/utils/guard"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

const (
	maxUserPasswordLength = 20
	minUserPasswordLength = 8
)

type Password struct {
	Value string
}

func NewPassword(value string) Password {
	return Password{
		Value: value,
	}
}

func NewPasswordWithValidation(value string, ctx context.Context) (Password, error) {
	err := guard.AgainstBetween("Password", value, minUserPasswordLength, maxUserPasswordLength)
	if err != nil {
		return Password{}, err
	}
	encryptedPassword, err := encryption.GenerateDefaultEncryption(
		value,
		env.GetEnvVariablesFromContext(ctx).Pepper,
	)
	if err != nil {
		return Password{}, web.InternalError()
	}
	return Password{Value: encryptedPassword}, nil
}
