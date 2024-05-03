package v1Router

import (
	"context"
	"fmt"
	"net/http"

	healthcheckUseCase "github.com/quinntas/go-rest-template/api/shared/useCases/healthCheck"
	userRouter "github.com/quinntas/go-rest-template/api/user/infra/http/router"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

func concatRouter(router web.HttpRouter, path string) string {
	return fmt.Sprintf("%s%s", router.Path, path)
}

func NewV1Router(path string, ctx context.Context) {
	router := web.NewHttpRouter(path, ctx)

	web.Route(router, http.MethodGet, "/health", ctx, healthcheckUseCase.UseCase)

	userRouter.NewUserRouter(concatRouter(router, "/users"), ctx)
}
