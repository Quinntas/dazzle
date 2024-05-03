package web

import (
	"fmt"
	"net/http"
)

func Serve(port string) {
	addr := fmt.Sprintf(":%s", port)

	notFoundHandler()

	err := http.ListenAndServe(addr, loggerHandler(http.DefaultServeMux))
	if err != nil {
		panic(err)
	}
}
