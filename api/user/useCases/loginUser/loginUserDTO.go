package loginUser

import (
	"context"

	"github.com/quinntas/go-rest-template/internal/api/web"
)

type DTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Token      string `json:"token"`
	ExpiresIn  int    `json:"expiresIn"`
	ExpireDate string `json:"expireDate"`
}

type PrivateTokenClaim struct {
	Id     int
	Email  string
	RoleId int
	UUID   string
}

type PublicTokenClaim struct {
	UUID string
}

func NewDTO(ctx context.Context) DTO {
	json := ctx.Value(web.JSON_CTX_KEY).(map[string]interface{})

	return DTO{
		Email:    json["email"].(string),
		Password: json["password"].(string),
	}
}
