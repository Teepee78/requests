package requests

import (
	"fmt"
	"net/http"
)

func (req *Request) Post(url string, data interface{}, config *Config) (*Response, error) {
	// parse request data
	parsedData, err := parseBody(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create request
	request, err := http.NewRequest(http.MethodPost, req.BaseUrl+url, parsedData)
	if err != nil {
		return nil, err
	}

	setHeaders(&request.Header, req.Headers, config.Headers)

	return sendRequest(request, config.Http)
}

func (req *Request) PostAsync(url string, body interface{}, config *Config) (chan *Response, error) {
	response := make(chan *Response)
	var err error

	go func() {
		res, getError := req.Post(url, body, config)
		response <- res
		err = getError
		close(response)
	}()

	return response, err
}
