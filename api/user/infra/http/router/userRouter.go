package userRouter

import (
	"context"
	"net/http"

	"github.com/quinntas/go-rest-template/api/user/useCases/createUser"
	"github.com/quinntas/go-rest-template/api/user/useCases/loginUser"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

func NewUserRouter(path string, ctx context.Context) {
	router := web.NewHttpRouter(path, ctx)

	web.Route(router, http.MethodPost, "/create", ctx, createUser.UseCase)
	web.Route(router, http.MethodPost, "/login", ctx, loginUser.UseCase)
}
