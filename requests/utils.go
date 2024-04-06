package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

/*
Parse request body to reader
*/
func parseBody(body interface{}) (io.Reader, error) {
	// Marshal the object into bytes
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Create an io.Reader from the JSON bytes
	reader := bytes.NewReader(jsonBytes)
	return reader, nil
}

func setHeaders(headers *http.Header, requestHeaders map[string]string, configHeaders map[string]string) {
	// set headers
	for key, value := range requestHeaders {
		headers.Set(key, value)
	}

	for key, value := range configHeaders {
		headers.Set(key, value)
	}
}

func sendRequest(request *http.Request, http bool) (*Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// read response body
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	// unmarshal response body
	var result interface{}
	if http {
		result = string(body)
	} else {
		json.Unmarshal(body, &result)
	}

	return &Response{
		Status:  response.StatusCode,
		Body:    result,
		Headers: response.Header,
		Request: response.Request,
		Cookies: response.Cookies(),
	}, nil
}
