package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func parseForm(request *http.Request) error {
	parseErr := request.ParseForm()
	if parseErr != nil {
		return BadRequest()
	}
	return nil
}

func parseQuery(rawQuery string) (url.Values, error) {
	queryValues, queryParserErr := url.ParseQuery(rawQuery)
	if queryParserErr != nil {
		return nil, BadRequest()
	}
	return queryValues, nil
}

func handleError(err error, response http.ResponseWriter) {
	if err != nil {
		switch t := err.(type) {
		case *HttpError:
			t.ToJsonResponse(response)
		}
	}
}

func HttpHandler(ctx context.Context, useCase UseCase) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		err := parseForm(request)
		if err != nil {
			handleError(err, response)
			return
		}

		query, err := parseQuery(request.URL.RawQuery)
		if err != nil {
			handleError(err, response)
			return
		}

		ctx = context.WithValue(ctx, QUERY_CTX_KEY, query)

		switch request.Header.Get("Content-Type") {
		case "application/json;charset=UTF-8":
		case "application/json; charset=UTF-8":
		case "application/json":
			var data interface{}
			err := json.NewDecoder(request.Body).Decode(&data)
			if err != nil {
				handleError(UnprocessableEntity(), response)
				return
			}
			ctx = context.WithValue(ctx, JSON_CTX_KEY, data)
		}

		err = useCase(request, response, ctx)
		if err != nil {
			handleError(err, response)
			return
		}
	}
}

func notFoundHandler() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		handleError(NotFound(), rw)
	})
}

func Route(router HttpRouter, method string, path string, ctx context.Context, useCase UseCase) {
	ctx = context.WithValue(ctx, PATH_CTX_KEY, path)
	http.HandleFunc(fmt.Sprintf("%s %s%s", method, router.Path, path), HttpHandler(ctx, useCase))
}
