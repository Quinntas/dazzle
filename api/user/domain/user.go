package userDomain

import (
	"context"

	sharedDomain "github.com/quinntas/go-rest-template/api/shared/domain"
	userValueObjects "github.com/quinntas/go-rest-template/api/user/domain/valueObjects"
	"github.com/quinntas/go-rest-template/internal/api/types"
)

type UserDomain struct {
	sharedDomain.SharedDomain
	Email    userValueObjects.Email
	Password userValueObjects.Password
	RoleId   types.ID
}

func NewUserDomainWithValidation(
	email string,
	password string,
	roleId int,
	ctx context.Context,
) (UserDomain, error) {
	emailValidated, err := userValueObjects.NewEmailWithValidation(email)
	if err != nil {
		return UserDomain{}, err
	}

	passwordValidated, err := userValueObjects.NewPasswordWithValidation(password, ctx)
	if err != nil {
		return UserDomain{}, err
	}

	return NewUserDomain(
		types.ID{},
		types.NewUUIDV4(),
		types.Date{},
		types.Date{},
		emailValidated,
		passwordValidated,
		types.NewID(roleId),
	), nil
}

func NewUserDomain(
	id types.ID,
	pid types.UUID,
	createdAt types.Date,
	updatedAt types.Date,
	email userValueObjects.Email,
	password userValueObjects.Password,
	roleId types.ID,
) UserDomain {
	return UserDomain{
		sharedDomain.NewSharedDomain(id, pid, createdAt, updatedAt),
		email,
		password,
		roleId,
	}
}
