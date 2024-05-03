package healthcheckUseCase

import (
	"context"
	"net/http"

	"github.com/quinntas/go-rest-template/internal/api/utils"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

func UseCase(request *http.Request, response http.ResponseWriter, ctx context.Context) error {
	web.JsonResponse(response, http.StatusOK, &utils.Map[string]{
		"message": "ok",
	})
	return nil
}
