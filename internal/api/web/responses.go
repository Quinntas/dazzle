package web

import (
	"encoding/json"
	"net/http"

	"github.com/quinntas/go-rest-template/internal/api/utils"
)

func DTOResponse[T interface{}](response http.ResponseWriter, status int, data *T) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	encoder := json.NewEncoder(response)
	if err := encoder.Encode(data); err != nil {
		// TODO
		return
	}
}

func JsonResponse[T interface{}](response http.ResponseWriter, status int, data *utils.Map[T]) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	encoder := json.NewEncoder(response)
	if err := encoder.Encode(data); err != nil {
		// TODO
		return
	}
}
