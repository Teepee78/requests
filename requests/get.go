package requests

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get(url string, config *Config) (*Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if config != nil && config.Headers != nil {
		for key, value := range config.Headers {
			request.Header.Set(key, value)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	var result interface{}
	if config != nil && config.Http {
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
