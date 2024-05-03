package sharedRouter

import (
	"context"

	v1Router "github.com/quinntas/go-rest-template/api/shared/infra/router/v1"
)

func InitRouters(ctx context.Context) {
	v1Router.NewV1Router("/api/v1", ctx)
}
