package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type ParsedBody[R interface{}] struct {
	Text  string
	Json  R
	Bytes []byte
	Type  string
}

type FetcherResult[R interface{}] struct {
	Status     int
	Ok         bool
	ParsedBody ParsedBody[R]
	Response   *http.Response
	TimeTaken  time.Duration
}

func (p *ParsedBody[any]) setText(body io.ReadCloser) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	p.Type = "text"
	p.Text = string(bodyBytes)
	return nil
}

func (p *ParsedBody[R]) setBytes(body io.ReadCloser) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	p.Type = "bytes"
	p.Bytes = bodyBytes
	return nil
}

func (p *ParsedBody[R]) setJson(body io.ReadCloser) error {
	err := json.NewDecoder(body).Decode(&p.Json)
	if err != nil {
		return err
	}
	p.Type = "json"
	return nil
}

func convertToBodyBytes[B interface{}](body *B, headers map[string]string) ([]byte, error) {
	var bodyBytes []byte
	var err error

	if headers["Content-Type"] == "" {
		return nil, errors.New("Content-Type header is required for non GET requests")
	}

	switch headers["Content-Type"] {
	case "application/json":
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		break
	// TODO: fix this shit
	case "application/ssml+xml":
		hJson, hErr := json.Marshal(body)
		if hErr != nil {
			return nil, err
		}
		var hData map[string]string
		hErr = json.Unmarshal(hJson, &hData)
		bodyBytes = []byte(hData["data"])
		break
	}

	return bodyBytes, nil
}

func Fetcher[B interface{}, R interface{}](method, url string, body *B, headers map[string]string) (*FetcherResult[R], error) {
	var req *http.Request
	var err error

	switch method {
	case http.MethodGet:
		req, _ = http.NewRequest(method, url, nil)
		break
	case http.MethodPost:
		bodyBytes, err := convertToBodyBytes[B](body, headers)
		if err != nil {
			return nil, err
		}
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
		break
	}

	req.Header.Set("User-Agent", "Encom/1.0.0")
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	startTimeTaken := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	endTimeTaken := time.Now()

	fetcherResult := FetcherResult[R]{
		Status:     resp.StatusCode,
		Ok:         resp.StatusCode >= 200 && resp.StatusCode < 300,
		Response:   resp,
		ParsedBody: ParsedBody[R]{},
		TimeTaken:  endTimeTaken.Sub(startTimeTaken),
	}

	switch resp.Header.Get("Content-Type") {
	case "application/json;charset=utf-8":
		err := fetcherResult.ParsedBody.setJson(resp.Body)
		if err != nil {
			return nil, err
		}
		break
	case "application/json":
		err := fetcherResult.ParsedBody.setJson(resp.Body)
		if err != nil {
			return nil, err
		}
		break
	case "application/json; charset=utf-8":
		err := fetcherResult.ParsedBody.setJson(resp.Body)
		if err != nil {
			return nil, err
		}
		break
	case "audio/webm; codec=opus":
		err := fetcherResult.ParsedBody.setBytes(resp.Body)
		if err != nil {
			return nil, err
		}
		break
	default:
		err := fetcherResult.ParsedBody.setText(resp.Body)
		if err != nil {
			return nil, err
		}
		break
	}

	return &fetcherResult, nil
}
