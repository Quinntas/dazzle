package web

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/quinntas/go-rest-template/internal/api/utils/terminal"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func printResponseTime(responseTime time.Duration) string {
	var colour string
	if responseTime.Seconds() > 1 {
		colour = terminal.Red
	} else if responseTime.Seconds() > 0.5 {
		colour = terminal.Yellow
	} else {
		colour = terminal.Green
	}
	return colour + responseTime.String() + terminal.Reset
}

func printStatus(status int) string {
	var colour string
	switch {
	case status >= 200 && status < 300:
		colour = terminal.Green
	case status >= 300 && status < 400:
		colour = terminal.Cyan
	case status >= 400:
		colour = terminal.Red
	default:
		colour = terminal.Reset
	}
	return colour + strconv.Itoa(status) + terminal.Reset
}

func printMethod(method string) string {
	var colour string
	switch method {
	case "GET":
		colour = terminal.Green
	case "POST":
		colour = terminal.Yellow
	case "PATCH":
		colour = terminal.Cyan
	case "DELETE":
		colour = terminal.Red
	default:
		colour = terminal.Reset
	}
	return colour + method + terminal.Reset
}

func (w *statusWriter) writeHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func loggerHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()
		applyDefaultHeaders(rw)
		sw := &statusWriter{ResponseWriter: rw}
		handler.ServeHTTP(sw, r)
		end := time.Now()
		responseTime := end.Sub(start)
		fmt.Println(r.RemoteAddr, printMethod(r.Method), r.URL, printResponseTime(responseTime), printStatus(sw.status))
	})
}
