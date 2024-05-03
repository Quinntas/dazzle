package loginUser

import (
	"context"
	"net/http"
	"time"

	userValueObjects "github.com/quinntas/go-rest-template/api/user/domain/valueObjects"
	userRepo "github.com/quinntas/go-rest-template/api/user/repo"
	"github.com/quinntas/go-rest-template/internal/api/encryption"
	"github.com/quinntas/go-rest-template/internal/api/redis"
	"github.com/quinntas/go-rest-template/internal/api/utils"
	"github.com/quinntas/go-rest-template/internal/api/utils/env"
	"github.com/quinntas/go-rest-template/internal/api/utils/jwtUtils"
	"github.com/quinntas/go-rest-template/internal/api/utils/timeUtils"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

func UseCase(request *http.Request, response http.ResponseWriter, ctx context.Context) error {
	dto := NewDTO(ctx)

	email, err := userValueObjects.NewEmailWithValidation(dto.Email)
	if err != nil {
		return err
	}

	_, err = userValueObjects.NewPasswordWithValidation(dto.Password, ctx)
	if err != nil {
		return err
	}

	user, err := userRepo.FindWithEmail(email.Value, ctx)
	if err != nil {
		return web.NewHttpError(
			http.StatusNotFound,
			"user not found",
			utils.Map[interface{}]{},
		)
	}

	env := env.GetEnvVariablesFromContext(ctx)

	if result, err := encryption.CompareEncryption(dto.Password, user.Password.Value, env.Pepper); err != nil || !result {
		return web.NewHttpError(
			http.StatusBadRequest,
			"credentials do not match",
			utils.Map[interface{}]{},
		)
	}

	redisClient := redis.GetRedisClientFromContext(ctx)

	privateToken, err := jwtUtils.GenerateJsonWebToken[PrivateTokenClaim](&PrivateTokenClaim{
		Id:     user.ID.Value,
		Email:  user.Email.Value,
		UUID:   user.PID.Value,
		RoleId: user.RoleId.Value,
	}, TokenExpirationTime, env.JwtSecret)
	if err != nil {
		return web.InternalError()
	}

	err = redisClient.Set(TokenRedisKey+user.PID.Value, privateToken, TokenExpirationTime)
	if err != nil {
		return web.InternalError()
	}

	publicToken, err := jwtUtils.GenerateJsonWebToken[PublicTokenClaim](&PublicTokenClaim{
		UUID: user.PID.Value,
	}, TokenExpirationTime, env.JwtSecret)
	if err != nil {
		return web.InternalError()
	}

	web.DTOResponse[ResponseDTO](response, http.StatusOK, &ResponseDTO{
		Token:      publicToken,
		ExpiresIn:  int(TokenExpirationTime.Seconds()),
		ExpireDate: timeUtils.TimeToString(time.Now().Add(TokenExpirationTime)),
	})

	return nil
}
