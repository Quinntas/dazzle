package web

import (
	"context"
	"net/http"
)

type UseCase func(request *http.Request, response http.ResponseWriter, ctx context.Context) error
