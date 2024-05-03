package createUser

import (
	"context"
	"net/http"

	userDomain "github.com/quinntas/go-rest-template/api/user/domain"
	userRepo "github.com/quinntas/go-rest-template/api/user/repo"
	"github.com/quinntas/go-rest-template/internal/api/utils"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

func UseCase(request *http.Request, response http.ResponseWriter, ctx context.Context) error {
	dto := NewDTO(ctx)

	user, err := userDomain.NewUserDomainWithValidation(
		dto.Email,
		dto.Password,
		USER_DEFAULT_ROLE_ID,
		ctx,
	)
	if err != nil {
		return err
	}

	_, err = userRepo.Create(user, ctx)
	if err != nil {
		return web.NewHttpError(
			http.StatusConflict,
			"Email already exists",
			utils.Map[interface{}]{},
		)
	}

	web.JsonResponse(response, http.StatusCreated, &utils.Map[string]{
		"message": "user created",
	})

	return nil
}
