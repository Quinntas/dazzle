package createUser

import (
	"context"

	"github.com/quinntas/go-rest-template/internal/api/web"
)

type DTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewDTO(ctx context.Context) DTO {
	json := ctx.Value(web.JSON_CTX_KEY).(map[string]interface{})

	return DTO{
		Email:    json["email"].(string),
		Password: json["password"].(string),
	}
}
