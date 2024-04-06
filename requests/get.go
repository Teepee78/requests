package requests

import (
	"encoding/json"
	"io"
	"net/http"
)

func (req *Request) Get(url string, config *Config) (*Response, error) {
	// create request
	request, err := http.NewRequest(http.MethodGet, req.BaseUrl+url, nil)
	if err != nil {
		return nil, err
	}

	// set headers
	for key, value := range req.Headers {
		request.Header.Set(key, value)
	}
	if config != nil && config.Headers != nil {
		for key, value := range config.Headers {
			request.Header.Set(key, value)
		}
	}

	// send request
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

func (req *Request) GetAsync(url string, config *Config) (chan *Response, error) {
	response := make(chan *Response)
	var err error

	go func() {
		res, getError := req.Get(url, config)
		response <- res
		err = getError
		close(response)
	}()

	return response, err
}
