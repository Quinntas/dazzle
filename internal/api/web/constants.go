package web

import (
	"net/http"

	"github.com/quinntas/go-rest-template/internal/api/utils"
)

const (
	JSON_CTX_KEY  utils.Key = "JSON"
	QUERY_CTX_KEY utils.Key = "QUERY"
	PATH_CTX_KEY  utils.Key = "PATH"
)

func applyDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Server", "Encom-Backend/1.0.0 (Golang)")
	w.Header().Set("X-Powered-By", "Encom")
}
